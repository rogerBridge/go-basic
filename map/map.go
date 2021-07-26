package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]int)
	var s sync.WaitGroup
	s.Add(2)
	go addMap(m, "age", 25, &s)
	go addMap(m, "height", 180, &s)
	s.Wait()
	fmt.Println(m)
}

func addMap(m map[string]int, key string, value int, sw *sync.WaitGroup){
	m[key] = value
	sw.Done()
}
// map类型的初始化方法
func mapInit() map[string]string {
	//kv := make(map[string]string, 100), 提前分配好map的空间可以防止运行时不断添加键值对, 提高程序的运行速度
	var kv map[string]string
	kv = map[string]string{}
	kv["name"] = "Fenr"
	kv["age"] = "24"
	return kv
}

// map类型的for-range用法
func mapForRange() {
	kv := make(map[string]string, 100)
	kv["name"] = "leo"
	kv["age"] = "24"
	kv["1"] = "1+1"
	kv["2"] = "2+2"
	kv["3"] = "3+3"
	for _, v := range kv {
		fmt.Println(v)
	}
}
