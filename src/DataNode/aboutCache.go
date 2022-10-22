package DataNode

//
//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//func FileReadStream(srcFile,maxCache,policy string)  {
//	fd, open := os.Open(srcFile)
//	if open != nil {
//		fmt.Println(open)
//	}
//	defer func() {
//		fd.Close()
//	}()
//	r := bufio.NewReader(fd)
//	var m []KV
//	readByte,_:=r.ReadBytes('\n')
//	json.Unmarshal(readByte, &m)
//	CacheHitRatio(m, maxCache, policy)
//}
