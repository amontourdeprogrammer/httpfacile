package main

import "fmt"
import "net"

func StartServer() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			// handle error
		}

		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, "Hello!")
		conn.CloseWrite()
	}
}

func main() {
	fmt.Printf("hello\n")
	StartServer()
}
