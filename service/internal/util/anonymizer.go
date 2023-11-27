package util

import "net"

func GetFakeClientIP() net.IP {
	return net.IPv4(127, 0, 0, 129)
}
