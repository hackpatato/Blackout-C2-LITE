package tcp

import (
	"bufio"
	"fmt"
	"net"
)


func Server() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("server not started", err)
		return
	}
	defer listener.Close()

	fmt.Println("tcp server is started , waiting connet...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error ", err)
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("new client connected", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading error:", err)
	}
	fmt.Println("client is disconnected:", conn.RemoteAddr().String())
}