package main

import (
	"fmt"
	"net"
)

func main() {
	ifaceName := "en0"
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		panic(err)
	}

	addr := &net.UDPAddr{
		IP:   net.ParseIP("ff02::1"),
		Port: 5000,
		Zone: iface.Name,
	}

	conn, err := net.ListenMulticastUDP("udp", iface, addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1024*1024)

	n, remoteAddr, _ := conn.ReadFromUDP(buffer)

	fmt.Println(string(buffer[:n]))
	fmt.Println(remoteAddr)
}
