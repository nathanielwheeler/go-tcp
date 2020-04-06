package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Spinning TCP mux...")

	// Initialize listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		// Accept connection requests
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Initial connection message
		io.WriteString(connection, "\nHello from TCP server!  Type something in.\n\n")

		go handle(connection)
	}
}

func handle(connection net.Conn) {
	// Initialize scanner loop
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		// Scan input into line
		line := scanner.Text()

		// Print response in mux log
		fmt.Println("Got message:", line)

		// Echo input to connection with formatting.
		fmt.Fprintf(connection, "I heard you say: \"%s\"\n", line)
	}
	defer connection.Close()
}
