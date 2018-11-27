package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Listening to connection on port 8080")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	log.Println("Accepted new connections")
	defer conn.Close()
	defer log.Println("Closed connection")

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		log.Println("Read new data from connection: ", data)
		conn.Write(data)
	}
}
