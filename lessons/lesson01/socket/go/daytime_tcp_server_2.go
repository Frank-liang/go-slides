package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime)) // don't care about return value
	conn.Close()                // we're finished with this client
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
