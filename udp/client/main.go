package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := net.UDPAddr{
		Port: 1200,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := "hello form the client"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server => %s ", buf[:n])
}
