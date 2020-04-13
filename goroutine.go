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

	//// example 2: multiple goroutine
	//go numbers()
	//go letters()
	//time.Sleep(3000 * time.Millisecond)
	//fmt.Println("main terminated")

	//// 猜想: 100个耗时2秒的非cpu任务, 如何不用waitGroup做到呢?
	//wg := new(sync.WaitGroup)
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go consumeTime(2, wg)
	//}
	//wg.wait()
	//fmt.Println("test finished")
	var count int = 0
	//m := new(sync.Mutex)
	m := make(chan bool)
	for i := 0; i < 100; i++ {
		go consumeTime(2, i, &count, m)
	}
	<-m
	for {
		if count == 100 {
			break
		}
	}
}

func consumeTime(i int, index int, count *int, m chan bool) {
	fmt.Printf("Sleep %d seconds, index is %d\n", i, index)
	time.Sleep(time.Duration(i) * time.Second)
	//m.Lock()
	m <- true
	*count += 1
	fmt.Printf("count value is %d\n", *count)
	<-m
	//m.Unlock()
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
