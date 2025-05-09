package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.TCPAddr{
		Port: 1200,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		tcpConn, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		n, err := tcpConn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("msg => %s \n", buf[:n])
		_, err = tcpConn.Write([]byte("HERE IS THE SERVER TCP ;) => YEAH the guy who was on the moon said this but i aslo don't rember the rest try anthoer tcp "))
	}

}
