package resolver

import (
	"fmt"
	"net"
)

func getIPRange(start, end string) ([]string, error) {
	startIP := net.ParseIP(start)
	endIP := net.ParseIP(end)

	if startIP == nil || endIP == nil {
		return nil, fmt.Errorf("invalid IP address")
	}

	startIP = startIP.To4()
	endIP = endIP.To4()
	if startIP == nil || endIP == nil {
		return nil, fmt.Errorf("IPv6 is not supported for output ranges:(")
	}

	// TODO: IPv6?

	startInt := ipToInt(startIP)
	endInt := ipToInt(endIP)

	var ips []string
	for i := startInt; i <= endInt; i++ {
		ips = append(ips, intToIP(i).String())
	}

	return ips, nil
}

func ipToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

func intToIP(n uint32) net.IP {
	return net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}
