package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(4)
	ch := make(chan int)
	go producer1(ch)
	go customer(ch)
	time.Sleep(time.Millisecond * 100)
}

func producer1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func customer(ch chan int) {
	for i := 0; ; i++ {
		fmt.Printf("%v ", <-ch)
	}
}

func sendData(ch chan string) {
	ch <- "安瑞峰"
	ch <- "木鱼"
	ch <- "渣滓"
	ch <- "菜鸡"
	ch <- "太难了呀"
}

func getData(ch chan string) []string {
	inputSlice := make([]string, 5)
	for i := 0; i < 5; i++ {
		inputSlice[i] = <-ch
	}
	fmt.Println(inputSlice)
	return inputSlice
}
