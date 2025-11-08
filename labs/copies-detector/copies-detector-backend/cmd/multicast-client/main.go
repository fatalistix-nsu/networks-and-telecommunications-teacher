package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	suitableInterfaces := make([]net.Interface, 0, len(interfaces))

	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 {
			continue
		}

		if i.Flags&net.FlagMulticast == 0 {
			continue
		}

		if i.Flags&net.FlagLoopback != 0 {
			continue
		}

		if i.Flags&net.FlagRunning == 0 {
			continue
		}

		if i.Flags&net.FlagUp == 0 {
			continue
		}

		addr, err := i.Addrs()
		if err != nil {
			continue
		}

		hasIpv6 := false
		for _, a := range addr {
			ip, _, err := net.ParseCIDR(a.String())
			if err != nil {
				continue
			}
			if ip.To4() != nil && ip.To16() == nil {
				hasIpv6 = true
				break
			}
		}
		if !hasIpv6 {
			continue
		}

		suitableInterfaces = append(suitableInterfaces, i)
	}

	fmt.Println(suitableInterfaces)

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

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
}
