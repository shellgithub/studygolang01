package main

import (
	"fmt"
	"net"
)

func main() {

	//监听
	listerner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listerner.Close()

	//阻塞等待用户链接
	for {

		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		//接收用户的请求
		buf := make([]byte, 1024) //1024大小的缓冲区
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err = ", err1)
			return
		}

		fmt.Println("buf = ", string(buf[:n]))

		defer conn.Close() //关闭当前用户链接

	}

}

//go run 01_TCP服务器.go

//nc 127.0.0.1 8000
// aaaaa
