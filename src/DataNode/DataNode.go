package DataNode

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/rpc"
)

// NodeInfo 用于json和结构体对象的互转
type NodeInfo struct {
	NodeName   string `json:"nodeName"`   //节点hostname 通过配置文件获取
	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
	Port       string `json:"port"`       //节点端口号
	//缓存列表 指针
	//命中率的信息
}

//DataNode节点的格式化输出
func (node NodeInfo) String() string {
	nn, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	return out.String()
}

// GetNode 获取Node
func GetNode(Name, Port, IP string) NodeInfo {
	NN := new(NodeInfo)
	NN.NodeName = Name
	NN.Port = Port
	NN.NodeIpAddr = IP
	return *NN
}

// StartDNClient 启动DataNode Client
func StartDNClient(dn NodeInfo, clusterIP string) {
	client, _ := rpc.DialHTTP("tcp", clusterIP)
	var nodelist []NodeInfo
	asyncCall := client.Go("NameNodeInfo.AddNode", dn, &nodelist, nil)
	<-asyncCall.Done
	for _, info := range nodelist {
		log.Printf(info.String() + " ")
	}
	log.Println()
	StartDNRPCServer(dn)
}
func StartDNRPCServer(dn NodeInfo) {
	rpc.Register(dn)
	rpc.HandleHTTP()

	log.Printf("Serving RPC server:%v on port %v", dn.NodeName, dn.Port)
	if err := http.ListenAndServe(":"+dn.Port, nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
