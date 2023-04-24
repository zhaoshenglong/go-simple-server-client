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

func main() {
	var (
		msg  = "Hello,world"
		res  = make([]byte, 1024)
		addr *net.TCPAddr
		conn net.Conn
		err  error
	)
	flag.Parse()
	if addr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", *host, *port)); err != nil {
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	util.EvalLatency("DialTCP", func() {
		if conn, err = net.DialTCP("tcp", nil, addr); err != nil {
			fmt.Println("Dial failed:", err.Error())
			os.Exit(1)
		}
	})

	util.EvalLatency("WriteMsg", func() {
		util.WriteMsg(conn, []byte(msg))
		fmt.Println("write to server = ", msg)
	})

	util.EvalLatency("ReadMsg", func() {
		res, _ = util.ReadMsg(conn)
		fmt.Println("reply from server=", string(res))
	})

	conn.Close()
}
