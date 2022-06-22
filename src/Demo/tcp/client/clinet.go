package client

import (
	"bufio"
	"fmt"
	"github/Gavinyjb/Gopkg/proto"
	"net"
	"os"
	"strings"
)

// NodeAsClient 客户端
func NodeAsClient(address string) {
	//address:="127.0.0.1:30000"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()                       // 关闭连接
	inputReader := bufio.NewReader(os.Stdin) // 读取用户输入
	for {
		msg, _ := inputReader.ReadString('\n')
		if strings.ToUpper(msg) == "Q" { // 如果输入q就退出
			return
		}
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		_, err = conn.Write(data)
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
