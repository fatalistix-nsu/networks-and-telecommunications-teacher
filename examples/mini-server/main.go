package main

import (
	"dev.gaijin.team/go/golib/must"
	"fmt"
	"net"
)

func main() {
	a := &net.TCPAddr{
		IP:   net.IPv4zero,
		Port: 8880,
	}

	l := must.OK(net.ListenTCP("tcp4", a))

	conn := must.OK(l.AcceptTCP())

	buffer := make([]byte, 10_000_000)

	read := must.OK(conn.Read(buffer))

	fmt.Println(read)
}
