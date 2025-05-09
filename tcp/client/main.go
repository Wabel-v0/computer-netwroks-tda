package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := net.TCPAddr{
		Port: 1200,
		IP:   net.ParseIP("127.0.01"),
	}
	conn, err := net.DialTCP("tcp", nil, &serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msg := "one setp for men ... don't rember the rest"
	_, err = conn.Write([]byte(msg))
	buf := make([]byte, 1024)
	res, _ := conn.Read(buf)
	fmt.Printf("%s", string(buf[:res]))
}
