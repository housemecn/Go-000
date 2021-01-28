package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: main
// @Version: 1.0.0
// @Date: 2021/1/28 23:02
// @Package Week09

// Message .
type Message struct {
	MsgChan chan string
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12339")
	if err != nil {
		log.Fatalf("listen error:%v\n", err)
	}
	fmt.Println("程序启动,开始监听12339端口...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		//启动读conn的协程
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	msg := &Message{make(chan string, 12)}
	go sendMsg(conn, msg.MsgChan)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error:%v\n", err)
			return
		}

		msg.MsgChan <- string(line)
	}
}

func sendMsg(conn net.Conn, ch <-chan string) {
	wr := bufio.NewWriter(conn)

	for msg := range ch {
		wr.WriteString("hello")
		wr.WriteString(msg)
		wr.Flush()
	}
}
