package server

import (
	"bufio"
	"fmt"
	"github/Gavinyjb/Gopkg/proto"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Printf("收到client:%v发来的数据：%v", conn, msg)
		fmt.Println("收到client发来的数据：", msg)
		//buf:=[512]byte{}
		conn.Write([]byte(msg)) //发送数据
	}
}

func NodeAsServer(address string) {

	//address:="127.0.0.1:30000"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理链接
	}
}
