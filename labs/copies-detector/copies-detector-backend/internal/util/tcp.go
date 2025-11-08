package util

import (
	"fmt"
	"net"
)

func ListenTCPAnyPort() (*net.TCPListener, int, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		return nil, 0, fmt.Errorf("could not resolve tcp address: %w", err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot listen tcp on addr %s: %w", addr, err)
	}

	tcpAddr, ok := l.Addr().(*net.TCPAddr)
	if !ok {
		panic("not a tcp address")
	}

	return l, tcpAddr.Port, nil
}
