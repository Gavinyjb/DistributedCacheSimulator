// server/main.go
package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}
type Strlist struct {
	Strlist []string
}

type Cal struct {
	StrList []string
}

func (c *Cal) Init(str string) {
	c.StrList = append(c.StrList, str)
}

func (cal *Cal) String(_ string, ret *[]string) error {
	*ret = make([]string, 0)
	*ret = cal.StrList
	return nil
}
func (cal *Cal) SetName(name string, ret *Strlist) error {
	cal.StrList = append(cal.StrList, name)
	for _, s := range cal.StrList {
		ret.Strlist = append(ret.Strlist, s)
	}
	//ret.Strlist=append(ret.Strlist,"yvjinbo")
	return nil
}
func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}
func (cal *Cal) Factorial(num int, result *Result) error {
	result.Num = num
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	result.Ans = sum
	return nil
}

func main() {
	demo := new(Cal)
	demo.Init("first")
	rpc.Register(demo)
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}

/*
使用 rpc.Register，发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
使用 rpc.HandleHTTP 注册用于处理 RPC 消息的 HTTP Handler
使用 http.ListenAndServe 监听 1234 端口，等待 RPC 请求。
*/
