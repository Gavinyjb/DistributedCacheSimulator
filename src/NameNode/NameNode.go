package NameNode

import (
	"../DataNode"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/rpc"
	"reflect"
	"sync"
)

var wg sync.WaitGroup

// NameNodeInfo 用于json和结构体对象的互转
type NameNodeInfo struct {
	NodeName   string              `json:"nodeName"`   //节点hostname 通过配置文件获取
	NodeIpAddr string              `json:"nodeIpAddr"` //节点ip地址
	Port       string              `json:"port"`       //节点端口号
	NodeList   []DataNode.NodeInfo `json:"NodeList"`   //存储已注册的DataNode节点
}

// SetNode 修改namenode节点信息
func (namenode *NameNodeInfo) SetNode(name string, ret *string) error {
	namenode.NodeName = name
	//ip 以及 端口号不可修改
	//namenode.NodeIpAddr = ip
	//namenode.Port = port
	*ret = fmt.Sprintf("namenode hostname已修改为%s\n", name)
	return nil
}
func SetNode(name, ip, port string) NameNodeInfo {
	namenode := new(NameNodeInfo)
	namenode.NodeName = name
	namenode.NodeIpAddr = ip
	namenode.Port = port
	return *namenode
}

//结构化输出字符串
func (namenode *NameNodeInfo) String(_ string, ret *string) error {
	nn, _ := json.Marshal(namenode)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	*ret = out.String()
	return nil
}

// GetNodeList 获取集群节点列表
func (namenode *NameNodeInfo) GetNodeList(_ string, ret *[]DataNode.NodeInfo) error {
	*ret = make([]DataNode.NodeInfo, 0)
	*ret = append(*ret, namenode.NodeList...)
	return nil
}

// AddNode 向集群中加入节点（注册节点）
func (namenode *NameNodeInfo) AddNode(newNode DataNode.NodeInfo, ret *[]DataNode.NodeInfo) error {
	namenode.NodeList = append(namenode.NodeList, newNode)
	*ret = make([]DataNode.NodeInfo, 0)
	*ret = namenode.NodeList
	return nil
}

// DelNode 从集群中移除节点
func (namenode *NameNodeInfo) DelNode(newNode DataNode.NodeInfo, ret *[]DataNode.NodeInfo) error {
	for i, node := range namenode.NodeList {
		if reflect.DeepEqual(node, newNode) {
			namenode.NodeList = append(namenode.NodeList[:i], namenode.NodeList[i+1:]...)
		}
	}
	*ret = make([]DataNode.NodeInfo, 0)
	*ret = namenode.NodeList
	return nil
}

//GetNodeIP 获取集群中某一节点的IP地址
func (namenode *NameNodeInfo) GetNodeIP(nodename string, ret *string) error {
	for _, node := range namenode.NodeList {
		if node.NodeName == nodename {
			*ret = node.NodeIpAddr + "|" + node.Port
			return nil
		}
	}
	return errors.New("集群中未查询到该节点！！！")
}

// StartNNRPCServer 启动NameNode RPC server
func StartNNRPCServer(nn NameNodeInfo) {
	namenode := new(NameNodeInfo)
	*namenode = nn
	rpc.Register(namenode)
	rpc.HandleHTTP()

	log.Printf("Serving RPC server:%v on port %v", nn.NodeName, nn.Port)
	if err := http.ListenAndServe(":"+nn.Port, nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}

func (namenode *NameNodeInfo) NN2DNClient(datanode DataNode.NodeInfo) {

}
