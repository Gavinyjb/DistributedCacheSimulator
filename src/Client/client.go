package Client

import (
	"../DataNode"
	"fmt"
	"net/rpc"
)

// StartClient 启动Client
func StartClient(message string, ServerIP string) {
	switch message {
	case "getDNList": //获取DataNode 列表
		client, _ := rpc.DialHTTP("tcp", ServerIP)
		var datanodelist []DataNode.NodeInfo
		asyncCall := client.Go("NameNodeInfo.GetNodeList", "", &datanodelist, nil)
		<-asyncCall.Done
		for _, info := range datanodelist {
			fmt.Printf(info.String() + " ")
		}
	case "GetDNInfo":
		client, _ := rpc.DialHTTP("tcp", ServerIP)
		fmt.Printf("获取%v节点信息\n", ServerIP)
		var dnInfo string
		asyncCall := client.Go("DataNodeInfo.GetDNInfo", "", &dnInfo, nil)
		<-asyncCall.Done
		fmt.Println(dnInfo)
	}
}
