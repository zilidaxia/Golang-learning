package main

import (
	"bufio"
	"log"
	"net"
)

// socks5协议不可翻墙 因为是明文传输

// TCP echo server （发啥回啥） 先用于测试server写的对否
func main() {
	// 监听端口 返回server
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept() //接受请求 如果成功返回连接
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client) //处理连接 go关键字 类似于java创建一个子线程，只是这里goroutine的开销会比子线程小很多
	}
}

// process函数实现
func process(conn net.Conn) {
	defer conn.Close()              //defer 表示这个函数退出的时候要把这个连接关掉，否则会有资源泄露
	reader := bufio.NewReader(conn) //带缓冲的只读流 减少底层系统调用次数，
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
