package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	demoTcpServer()
}

func demoTcpServer() {
	tcpAddr1 := net.TCPAddr{
		IP:   nil,
		Port: 9090,
		Zone: "",
	}
	listen1, err := net.ListenTCP("tcp", &tcpAddr1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen1.AcceptTCP()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		println(conn.RemoteAddr(), "连接")
		openKeepAliveErr := conn.SetKeepAlive(true)
		if openKeepAliveErr != nil {
			fmt.Println(openKeepAliveErr)
			return
		}
		setErr := conn.SetKeepAlivePeriod(5 * time.Second)
		if setErr != nil {
			fmt.Println(setErr)
			return
		}
		//create goroutine for each connect
		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
