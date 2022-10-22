package Client

import (
	"../DataNode"
	"fmt"
	"net/rpc"
)

// StartClient 启动Client
func StartClient(message string, NameNodeIP string) {
	client, _ := rpc.DialHTTP("tcp", NameNodeIP)
	switch message {
	case "getDNList": //获取DataNode 列表
		var datanodelist []DataNode.NodeInfo
		asyncCall := client.Go("NameNodeInfo.GetNodeList", "", &datanodelist, nil)
		<-asyncCall.Done
		for _, info := range datanodelist {
			fmt.Printf(info.String() + " ")
		}
	case "":

	}
}
