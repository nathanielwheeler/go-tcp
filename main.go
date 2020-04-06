package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello from go-tcp.main!")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(connection, "\nHello from TCP server!\n")

		connection.Close()
	}
}
