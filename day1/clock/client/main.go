package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp","localhost:13301")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout,conn)
}

func mustCopy(dst *os.File, src net.Conn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
