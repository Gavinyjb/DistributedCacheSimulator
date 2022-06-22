package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"
)

// NodeInfo 用于json和结构体对象的互转
type NodeInfo struct {
	NodeName   string `json:"nodeName"`   //节点hostname
	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
	Port       string `json:"port"`       //节点端口号
}
//DataNode节点的格式化输出
func (node NodeInfo) string()string  {
	nn,_:=json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out,nn,"","\t")
	return  out.String()
}
// InitDN DataNode节点初始化
func InitDN(NodeName,NodeIpAddr,Port string) NodeInfo {
	return NodeInfo{
		NodeName:   NodeName,
		NodeIpAddr: NodeIpAddr,
		Port:       Port,
	}
}
// NameNodeInfo 用于json和结构体对象的互转
type NameNodeInfo struct {
	NodeName   string     `json:"nodeName"`   //节点hostname
	NodeIpAddr string     `json:"nodeIpAddr"` //节点ip地址
	Port       string     `json:"port"`       //节点端口号
	NodeList   []NodeInfo `json:"NodeList"`   //存储已注册的DataNode节点
}
//NameNode节点信息的格式化输出
func (namenode NameNodeInfo) string() string {
	nn, _ := json.Marshal(namenode)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	return out.String()
}

// InitNN NameNode节点初始化
func InitNN(NodeName,NodeIpAddr,Port string) NameNodeInfo {
	return NameNodeInfo{
		NodeName:   NodeName,
		NodeIpAddr: NodeIpAddr,
		Port:       Port,
		NodeList:   nil,
	}
}
// MessageInfo 节点间发起一个请求或者响应的标准格式（节点间通讯信息）
type MessageInfo struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}
//将节点间通讯信息格式化
func (msg MessageInfo) String() string {
	nn, _ := json.Marshal(msg)
	var out bytes.Buffer
	json.Indent(&out, nn, "", "\t")
	return out.String()
}

func main() {
	//节点类型参数
	nodeType:=flag.String("nodeType","NameNode","请输入节点类型：NameNode,DataNode,Client")
	clusterIp := flag.String("clusterIp", "127.0.0.1:30000", "ip address of any node to connect")
	myPort := flag.String("myPort", "30000", "ip address to run this node on. default is 30000.")
	myName := flag.String("myName", "master", "node hostname")
	flag.Parse()

	//获取ip地址
	myIp, _ := net.InterfaceAddrs()

	//创建Node结构体
	switch *nodeType {
	case "NameNode":
		//创建NameNode结构体
		me:=InitNN(*myName,myIp[0].String(),*myPort)
		fmt.Println("我的节点信息：",me.string())
		//启动NN节点
		fmt.Println("将启动me节点为NameNode节点")
		StartNN(me)
	case "DataNode":
		//创建DataNode结构体
		me:=InitDN(*myName,myIp[0].String(),*myPort)
		fmt.Println("我的节点信息：",me.string())
		//启动DN节点
		fmt.Println("将启动me节点为DataNode节点")
		StartDN(me,*clusterIp)
	}
}

// StartNN 启动NameNode节点
func StartNN(me NameNodeInfo)  {
	//监听即将到来的信息
	listen, _:= net.Listen("tcp",":"+me.Port )
	defer listen.Close()
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go me.processNN(conn)//启动一个goroutine处理连接
	}
}

func (NN *NameNodeInfo)processNN(conn net.Conn) {
	defer conn.Close()
	//初次连接，注册节点
	var datanode NodeInfo
	json.NewDecoder(conn).Decode(&datanode)
	NN.NodeList = append(NN.NodeList,datanode )
	json.NewEncoder(conn).Encode(&NN)
	for {
		var msg NodeInfo
		err := json.NewDecoder(conn).Decode(&msg)
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Printf("收到client:%v发来的数据：%v",conn,msg.string())
		json.NewEncoder(conn).Encode(msg)
	}
}

// StartDN 启动DateNode节点
func StartDN(me NodeInfo,clusterIp string) bool {
	conn, err := net.DialTimeout("tcp", clusterIp,time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("不能连接到集群", me.NodeName)
			return false
		}
	}
	defer conn.Close()  // 关闭连接
	for i := 0; i < 2; i++ {
		json.NewEncoder(conn).Encode(&me)

		decoder := json.NewDecoder(conn)
		var responseMessage NameNodeInfo
		decoder.Decode(&responseMessage) //响应信息
		fmt.Println("得到数据响应:\n" + responseMessage.string())
	}
	return true
	//for  {
	//	json.NewEncoder(conn).Encode(&me)
	//
	//	decoder := json.NewDecoder(conn)
	//	var responseMessage NameNodeInfo
	//	decoder.Decode(&responseMessage) //响应信息
	//	fmt.Println("得到数据响应:\n" + responseMessage.string())
	//}
}
