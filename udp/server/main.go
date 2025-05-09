package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	addr := net.UDPAddr{
		Port: 1200,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		fmt.Println("Server is running")
		n, clientAdrr, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		msg := string(buf[:n])
		fmt.Printf("got: %s form %s\n", msg, clientAdrr)

		reply := []byte(strings.ToUpper(msg))
		_, err = conn.WriteToUDP(reply, clientAdrr)
		if err != nil {
			panic(err)
		}
	}

}
