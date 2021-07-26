package main

import (
	"fmt"
	"time"
)

func printNum(a int) {
	fmt.Println(a)
}

func foo(flag chan int) {
	time.Sleep(1*time.Second)
	flag <- 1
}

func main() {
	flag := make(chan int)
	go foo(flag)
	<-flag
}
