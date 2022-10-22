package main

import (
	"./Client"
	"./DataNode"
	"./NameNode"
	"flag"
	"fmt"
	"net"
	"sync"
)

//// NodeInfo 用于json和结构体对象的互转
//type NodeInfo struct {
//	NodeName   string `json:"nodeName"`   //节点hostname 通过配置文件获取
//	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
//	Port       string `json:"port"`       //节点端口号
//}
//
//// NameNodeInfo 用于json和结构体对象的互转
//type NameNodeInfo struct {
//	NodeName   string `json:"nodeName"`   //节点hostname 通过配置文件获取
//	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
//	Port       string `json:"port"`       //节点端口号
//	NodeList   []NodeInfo  `json:"NodeList"`  //存储已注册的DataNode节点
//}
var wg sync.WaitGroup

func main() {
	//节点类型参数
	nodeType := flag.String("nodeType", "NameNode", "请输入节点类型：NameNode,DataNode,Client")
	clusterIp := flag.String("clusterIp", "127.0.0.1:30000", "ip address of any node to connect")
	myPort := flag.String("myPort", "30000", "ip address to run this node on. default is 30000.")
	myName := flag.String("myName", "master", "node hostname")
	command := flag.String("command", "getDNList", "客户端传输指令：getDNList；")
	flag.Parse()

	//获取ip地址
	myIp, _ := net.InterfaceAddrs()

	switch *nodeType {
	case "NameNode":
		//启动NN节点
		fmt.Println("将启动me节点为NameNode节点")

		nn := NameNode.SetNode(*myName, myIp[0].String(), *myPort)
		wg.Add(1)
		go NameNode.StartNNRPCServer(nn)
		wg.Wait()
	case "DataNode":
		//启动DN节点
		fmt.Println("将启动me节点为DataNode节点")
		//初始化datanode 并返回结构体
		dn := DataNode.GetNode(*myName, myIp[0].String(), *myPort)
		client := DataNode.StartDNClient(dn, *clusterIp)
		fmt.Println(DataNode.GetNodeIP("slave4", client))
	case "Client":
		//启动客户端
		fmt.Println("将启动me节点为Client节点")
		Client.StartClient(*command, *clusterIp)
	}

}
