package multicast

import (
	"encoding/binary"
	"fmt"
	"github.com/fatalistix/slogattr"
	"log/slog"
	"net"
	"strconv"
	"sync"
	"time"
)

type Sender struct {
	log          *slog.Logger
	host         string
	port         int
	name         string
	sendInterval time.Duration
	shutdownChan chan bool
	shutdownWg   *sync.WaitGroup
}

func NewSender(
	log *slog.Logger,
	host string,
	port int,
	name string,
	sendInterval time.Duration,
) *Sender {
	shutdownWg := &sync.WaitGroup{}
	shutdownWg.Add(1)

	return &Sender{
		log:          log,
		host:         host,
		port:         port,
		name:         name,
		sendInterval: sendInterval,
		shutdownChan: make(chan bool),
		shutdownWg:   shutdownWg,
	}
}

func (s *Sender) Run() error {
	defer s.shutdownWg.Done()

	portStr := strconv.Itoa(s.port)
	hostPort := net.JoinHostPort(s.host, portStr)

	iface, err := net.InterfaceByName("en0")
	if err != nil {
		return fmt.Errorf("could not find en0 interface: %w", err)
	}

	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(s.host),
		Port: s.port,
		Zone: iface.Name,
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return fmt.Errorf("dial %s: %v", udpAddr, err)
	}

	defer s.closeConn(conn)

	joinOrRefreshMessage := make([]byte, typeSize+nameLengthSize+len(s.name))

	joinOrRefreshMessage[0] = byte(JoinOrRefresh)
	binary.BigEndian.PutUint16(joinOrRefreshMessage[1:3], uint16(len(s.name)))
	copy(joinOrRefreshMessage[3:], s.name)

	leaveMessage := make([]byte, typeSize+nameLengthSize+len(s.name))

	leaveMessage[0] = byte(Leave)
	binary.BigEndian.PutUint16(leaveMessage[1:3], uint16(len(s.name)))
	copy(leaveMessage[3:], s.name)

	t := time.NewTicker(s.sendInterval)

	for {
		select {
		case <-s.shutdownChan:
			s.log.Info("Sending leave message")
			_, err = conn.Write(leaveMessage)
			if err != nil {
				s.log.Warn("Sending leave message failed", slogattr.Err(err))
				return fmt.Errorf("write to %s: %w", hostPort, err)
			} else {
				s.log.Info("Stopping multicast sender")
				return nil
			}
		case <-t.C:
			_, err = conn.Write(joinOrRefreshMessage)
			if err != nil {
				return fmt.Errorf("write to %s: %w", hostPort, err)
			}
		}
	}
}

func (s *Sender) closeConn(conn *net.UDPConn) {
	err := conn.Close()
	if err != nil {
		s.log.Warn("close connection", slogattr.Err(err))
	}
}

func (s *Sender) Shutdown() {
	s.shutdownChan <- true
	s.shutdownWg.Wait()
}
