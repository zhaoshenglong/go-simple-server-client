package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"tcp-example/util"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 5001, "The port to connect to; defaults to 5001.")
var size = flag.Int("size", 1024, "The size of message to be sent; defaults to 1024.")

func main() {
	var (
		msg  string
		res  []byte
		addr *net.UDPAddr
		conn *net.UDPConn
		err  error
	)
	// Resolve the string address to a UDP address
	if addr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", *host, *port)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	util.EvalLatency("DialUDP", func() {
		if conn, err = net.DialUDP("udp", nil, addr); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})
	//
	// Send a message to the server
	msg = util.RandStringRunes(*size)

	util.EvalLatency("Write", func() {
		if _, err = conn.Write([]byte(msg)); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	util.EvalLatency("Read", func() {
		if _, err = conn.Read(res); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

}
