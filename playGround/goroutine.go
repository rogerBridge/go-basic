package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t0 := time.Now()
	flag := make(chan bool)
	concurrentNum := 10000
	ch := make(chan int, concurrentNum)
	go pump1(ch, concurrentNum)
	go suck1(ch, concurrentNum, flag)
	// 当suck函数处理完毕时, 取消阻塞
	<-flag
	fmt.Println(time.Since(t0))
}

func pump1(ch chan int, concurrentNum int)  {
	for i:=0; i<concurrentNum; i++{
		go randResult(ch)
	}
}

func suck1(ch chan int, concurrentNum int, flag chan bool) {
	resultSlice := make([]int, concurrentNum)
	// 确认从通道取值的数量, 可以让下面的代码reachable
	for i:=0; i<concurrentNum; i++{
		resultSlice[i] = <-ch
	}
	fmt.Println(float64(sumSlice(resultSlice))/float64(concurrentNum))
	flag<-true
	//var input int
	//for {
	//	input = <-ch
	//	fmt.Printf("%d", input)
	//}
}

func randResult(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	ch <- r
}

func sumSlice(s []int) int {
	sliceLength := len(s)
	sum := 0
	for i:=0; i<sliceLength; i++{
		sum += s[i]
	}
	return sum
}