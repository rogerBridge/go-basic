package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; i<1000000; i++ {
		ch <- i
	}
}

// 给suck函数划定有限的范围
func suck(ch chan int, flag chan bool) {
	t0 := time.Now()
	for i:=0; i<1000000; i++{
		fmt.Println(<-ch)
	}
	t1 := time.Since(t0)
	fmt.Println(t1)
	flag <- true
}

func main() {
	f := make(chan bool)
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1, f)
	<-f
}
