package main

import (
	"fmt"
	"time"
	_ "time"
)

func main() {
	c := make(chan int, 10)
	// 如果把c<-0放到前面的话, 会直接阻塞goroutine导致退出
	go func() {
		time.Sleep(10*time.Second)
		<-c
	}()
	c<-10
}
