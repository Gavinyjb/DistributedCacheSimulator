package main

import (
	"./DataNode"
	"./NameNode"
	"flag"
	"fmt"
	"net"
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
func main() {
	//节点类型参数
	nodeType:=flag.String("nodeType","NameNode","请输入节点类型：NameNode,DataNode,Client")
	clusterIp := flag.String("clusterIp", "127.0.0.1:30000", "ip address of any node to connect")
	myPort := flag.String("myPort", "30000", "ip address to run this node on. default is 30000.")
	myName := flag.String("myName", "master", "node hostname")
	flag.Parse()

	//获取ip地址
	myIp, _ := net.InterfaceAddrs()

	switch *nodeType {
	case "NameNode":
		//启动NN节点
		fmt.Println("将启动me节点为NameNode节点")

		nn:=DataNode.GetNode(*myName,*myPort,myIp[0].String())
		//nn:=new(NameNode.NodeInfo)
		//nn.NodeName=*myName
		//nn.Port=*myPort
		//nn.NodeIpAddr=myIp[0].String()
		NameNode.StartNNRPCServer(nn)
	case "DataNode":
		//启动DN节点
		fmt.Println("将启动me节点为DataNode节点")
		nn:=DataNode.GetNode(*myName,*myPort,myIp[0].String())
		DataNode.StartDNClient(nn,*clusterIp)
	}

}