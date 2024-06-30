//Writes the current system time to a client

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func _main() {
	portPtr := flag.Int("port", int(8000), "The port number to run")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portPtr))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Connection aborted
			continue
		}
		go handleConn(conn) // Handle via a goroutine
	}
}

func _handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // Client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
