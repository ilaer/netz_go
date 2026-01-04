package util

import (
	"fmt"
	"net"
)

func GetLocalIP(addr string) (ip string, err error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		fmt.Printf("无法连接到 %s: %v\n", addr, err)
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = localAddr.IP.String()

	return ip, nil
}
