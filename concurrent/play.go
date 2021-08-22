package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan int)
	// 如果把c<-0放到前面的话, 会直接阻塞goroutine导致退出
	go func() {
		log.Println("time sleep 10 seconds")
		time.Sleep(10 * time.Second)
		c<-0
	}()
	<-c
}
