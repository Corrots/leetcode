package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	Server()
}

func Server() {
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err: %v\n", err)
			continue
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		_, err := reader.Read(buf[:])
		if err != nil {
			log.Fatalf("read from conn err: %v\n", err)
		}
		fmt.Printf("received msg from: %s -> %v\n", conn.LocalAddr(), string(buf[:]))
		// 向客户端返回信息

		msg := "time: " + time.Now().Format("2006-01-02 15:04:05")
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatalf("write from conn err: %v\n", err)
		}
	}
}
