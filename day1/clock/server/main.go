package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":13301")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go Handle(conn)
	}
}

func Handle(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			// 如果发生错误，返回，关闭连接
			return
		}
		time.Sleep(1 * time.Second)
	}
}
