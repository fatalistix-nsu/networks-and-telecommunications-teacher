package multicast

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/fatalistix/slogattr"
)

type Handler func(addr string, name string)

type Listener struct {
	log          *slog.Logger
	host         string
	port         int
	handlers     map[MessageType]Handler
	shutdownChan chan bool
	inShutdown   *atomic.Bool
	shutdownWg   *sync.WaitGroup
}

func NewListener(
	log *slog.Logger,
	host string,
	port int,
) *Listener {
	shutdownWg := &sync.WaitGroup{}
	shutdownWg.Add(1)

	return &Listener{
		log:          log,
		host:         host,
		port:         port,
		handlers:     make(map[MessageType]Handler),
		shutdownChan: make(chan bool),
		inShutdown:   &atomic.Bool{},
		shutdownWg:   shutdownWg,
	}
}

func (l *Listener) SetHandler(t MessageType, handler Handler) {
	switch t {
	case JoinOrRefresh:
		l.handlers[JoinOrRefresh] = handler
	case Leave:
		l.handlers[Leave] = handler
	default:
		panic("unexpected message type" + strconv.Itoa(int(t)))
	}
}

func (l *Listener) Run() error {
	defer l.shutdownWg.Done()

	portStr := strconv.Itoa(l.port)
	hostPort := net.JoinHostPort(l.host, portStr)

	iface, err := net.InterfaceByName("en0")
	if err != nil {
		return fmt.Errorf("could not find en0 interface: %w", err)
	}

	addr := &net.UDPAddr{
		IP:   net.ParseIP(l.host),
		Port: l.port,
		Zone: iface.Name,
	}

	conn, err := net.ListenMulticastUDP("udp", iface, addr)
	if err != nil {
		return fmt.Errorf("listen multicast udp %s: %w", hostPort, err)
	}

	defer l.closeConn(conn)

	buf := make([]byte, typeSize+nameLengthSize+maxNameLength)

	go func() {
		<-l.shutdownChan
		l.inShutdown.Store(true)
		l.closeConn(conn)
	}()

	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			if l.shuttingDown() {
				l.log.Info("Listener is shutting down")
				return nil
			} else {
				return fmt.Errorf("read from %s: %w", hostPort, err)
			}
		}

		if n < typeSize+nameLengthSize {
			l.log.Warn(
				"Message too short",
				slog.Int("message_length", n),
				slog.Any("message", buf[:n]),
				slog.String("from", addr.String()),
			)
			continue
		}

		var msgType MessageType
		switch MessageType(buf[0]) {
		case JoinOrRefresh:
			msgType = JoinOrRefresh
		case Leave:
			msgType = Leave
		default:
			l.log.Warn(
				"Unsupported message type",
				slog.Int("message_type", int(msgType)),
				slog.Any("message", buf[:n]),
				slog.String("from", addr.String()),
			)
			continue
		}

		nameLen := int(binary.BigEndian.Uint16(buf[typeSize : typeSize+nameLengthSize]))
		if nameLen < 0 || maxNameLength < nameLen {
			nameLenLittleEndian := int(binary.LittleEndian.Uint16(buf[typeSize : typeSize+nameLengthSize]))
			l.log.Warn(
				"Invalid name length",
				slog.Int("name_length", nameLen),
				slog.Int("name_length_little_endian", nameLenLittleEndian),
				slog.Any("message", buf[:n]),
				slog.String("from", addr.String()),
			)
			continue
		}

		payload := buf[typeSize+nameLengthSize : typeSize+nameLengthSize+nameLen]

		h, ok := l.handlers[msgType]
		if !ok {
			return fmt.Errorf("no handler for message type: %s", strconv.Itoa(int(msgType)))
		}

		h(addr.String(), string(payload))
	}
}

func (l *Listener) shuttingDown() bool {
	return l.inShutdown.Load()
}

func (l *Listener) closeConn(conn *net.UDPConn) {
	err := conn.Close()
	if err != nil {
		if !errors.Is(err, net.ErrClosed) {
			l.log.Error("close connection", slogattr.Err(err))
		}
	}
}

func (l *Listener) Shutdown() {
	l.shutdownChan <- true
	l.shutdownWg.Wait()
}
