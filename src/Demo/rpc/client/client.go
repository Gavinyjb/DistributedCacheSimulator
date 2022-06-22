// client/main.go
package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
	Str      string
}
type Strlist struct {
	Strlist []string
}

//同步
//func main() {
//	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
//	var result Result
//	if err := client.Call("Cal.Square", 12, &result); err != nil {
//		log.Fatal("Failed to call Cal.Square. ", err)
//	}
//	log.Printf("%d^2 = %d", result.Num, result.Ans)
//}
//异步
func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	client.Call("Cal.Factorial", 5, &result)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
	var strList2 []string
	client.Call("Cal.String", "", &strList2)
	log.Println(strList2)
	var strList Strlist
	client.Call("Cal.SetName", "yvjinbo", &strList)
	log.Println(strList)

	client.Call("Cal.String", "", &strList2)
	log.Println(strList2)
}

/*
在客户端的实现中，因为要用到 Result 类型，简单起见，我们拷贝了 Result 的定义。

使用 rpc.DialHTTP 创建了 HTTP 客户端 client，并且创建了与 localhost:1234 的链接，1234 恰好是 RPC 服务监听的端口。
使用 rpc.Call 调用远程方法，第1个参数是方法名 Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。
*/
