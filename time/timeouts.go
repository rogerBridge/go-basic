package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	flag := make(chan bool, 1) // 建立一个flag, 用来保存程序执行之后的结果
	go func() {
		time.Sleep(1 * time.Second) // 用来模拟定时器执行之后的结果
		flag <- true
	}()
	select {
	// select 里面的chan是竞争关系, 谁先得到结果就选谁
	case res := <-flag:
		fmt.Printf("%+v %T %v\n", res, res, reflect.TypeOf(res).String())
	case <-time.After(5*time.Second):
		fmt.Println("time out")
	}
}