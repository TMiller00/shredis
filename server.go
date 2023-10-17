package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	for {
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	resp := NewResp(conn)
	value, err := resp.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)

	// ignore request and send back a PONG
	conn.Write([]byte("+PONG\r\n"))
}
