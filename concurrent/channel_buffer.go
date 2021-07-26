package main

import (
	"fmt"
	"time"
)

func main() {
	// 一个无缓冲的chan, 放入和取出必须一起做
	c := make(chan int)
	go func() {
		time.Sleep(10 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	// 填充50个才会阻塞
	c <- 10
	fmt.Println("sent", 10)
}
