package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理程序
func worker(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte("hello world"))
		break
	}
}

func main() {
	// 监听端口
	fmt.Println("listen: 9000")
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	// 接收请求
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go worker(conn)
	}
}
