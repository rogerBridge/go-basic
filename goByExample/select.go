package main

import (
	"fmt"
	"time"
)

func main() {
	c0 := make(chan int)
	c1 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		c0 <- 0
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- 1
	}()
	select { // select 经常用于超时处理
	case msg0 := <-c0:
		fmt.Println(msg0)
	case msg1 := <-c1:
		fmt.Println(msg1)
	}
}
