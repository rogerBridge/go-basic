package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine :)")
}

func playGoroutine() {
	//// example 1
	//// main goroutine 不会等到 hello goroutine执行完毕
	//// 可以理解为一个时间轴内, main goroutine 和 hello goroutine先后开始, 同时执行, main goroutine先结束, 整个程序没了... hello goroutine的结果也不能输出
	//go hello()
	//// time.Sleep(18*time.Microsecond)
	//fmt.Println("Hello")

	// example 2: multiple goroutine
	go numbers()
	go letters()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")

}

func numbers() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func letters() {
	for i := 'a'; i < 'f'; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c ", i)
		time.Sleep(30 * time.Millisecond)
	}
}
