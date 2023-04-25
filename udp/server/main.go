package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var port = flag.Int("port", 5001, "The port to connect to; defaults to 5001.")

func main() {
	var (
		addr *net.UDPAddr
		conn *net.UDPConn
		err  error
	)

	// Resolve the string address to a UDP address
	if addr, err = net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", *port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Listening at :%d\n", *port)

	// Start listening for UDP packages on the given address
	if conn, err = net.ListenUDP("udp", addr); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from UDP listener in endless loop
	for {
		var buf = make([]byte, 8192)
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Receive msg from %s, %s\n", addr.String(), string(buf))
		// Write back the message over UPD
		conn.WriteToUDP(buf, addr)
	}
}
