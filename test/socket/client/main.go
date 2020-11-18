package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	Client()
}

func Client() {
	conn, err := net.Dial("tcp", ":6666")
	if err != nil {
		log.Fatalf("conn server err: %v\n", err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err := conn.Write([]byte(s))
		if err != nil {
			log.Fatalf("send data from conn err: %v\n", err)
		}
		//
		var buf [1024]byte
		if _, err := conn.Read(buf[:]); err != nil {
			log.Fatalf("read data from conn err : %v\n", err)
		}
		fmt.Printf("收到服务端回复:%v\n", string(buf[:]))
	}
}
