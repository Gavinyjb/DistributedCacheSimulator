package analyer

import (
	"../../../../GolandProjects/DistributedCacheSimulator/src/policy/clock"
	"../../../../GolandProjects/DistributedCacheSimulator/src/policy/fifo"
	"../../../../GolandProjects/DistributedCacheSimulator/src/policy/lfu"
	"../../../../GolandProjects/DistributedCacheSimulator/src/policy/lru"
	"../../../../GolandProjects/DistributedCacheSimulator/src/policy/mru"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type kv struct {
	BlockId string `json:"块ID"`
	DestIp  string `json:"目的IP"`
}

func Analyer(srcFile, maxCaches, policies string) {
	fd, open := os.Open(srcFile)
	if open != nil {
		fmt.Println(open)
	}
	defer func() {
		fd.Close()
	}()
	r := bufio.NewReader(fd)
	var m []kv
	readString, _ := r.ReadString('\n')
	json.Unmarshal([]byte(readString), &m)
	maxCacheList := strings.Split(maxCaches, ",")
	policylist := strings.Split(policies, ",")
	for _, v := range maxCacheList {
		for _, policy := range policylist {
			maxCache, _ := strconv.ParseInt(v, 10, 64)
			CacheHitRatio(m, maxCache, policy)
		}
	}
}

func CacheHitRatio(m []kv, maxCache int64, policy string) {
	switch policy {
	case "lru":
		cache := lru.NewCache[string, string](lru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "lfu":
		cache := lfu.NewCache[string, string](lfu.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "mru":
		cache := mru.NewCache[string, string](mru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "fifo":
		cache := fifo.NewCache[string, string](fifo.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	case "clock":
		cache := clock.NewCache[string, string](clock.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	default:
		//默认LRU
		cache := lru.NewCache[string, string](lru.WithCapacity(int(maxCache)))
		hit, unhit := 0, 0
		for _, kv := range m {
			_, ok := cache.Get(kv.BlockId)
			if ok {
				//fmt.Println("命中-----" + got)
				hit++
			} else {
				//fmt.Println("未命中----" + got)
				unhit++
				cache.Set(kv.BlockId, kv.DestIp)
			}
		}
		var res float64
		res = float64(hit) / float64(hit+unhit)
		res *= 100
		fmt.Printf("当前容量：%v;当前置换算法的:%v;命中率：%v%%\n", maxCache, policy, res)
	}

}
