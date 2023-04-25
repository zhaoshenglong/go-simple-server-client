package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"tcp-example/util"
)

var port = flag.Int("port", 5001, "The port to connect to; defaults to 5001.")

// main serves as the program entry point
func main() {
	flag.Parse()
	// create a tcp listener on the given port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s\n", listener.Addr())

	// listen for new connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		// pass an accepted connection to a handler goroutine
		go handleConnection(conn)
	}
}

// handleConnection handles the lifetime of a connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		var (
			b   []byte
			err error
		)
		if b, err = util.ReadMsg(conn); err != nil {
			break
		}
		fmt.Printf("request: %s", b)
		if err = util.WriteMsg(conn, b); err != nil {
			break
		}
	}
}
