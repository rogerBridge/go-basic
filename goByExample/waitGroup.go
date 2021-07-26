package main

import (
	"sync"
	"time"
)

// waitGroup 可以用来等待多个goroutine完成

func worker1(wg *sync.WaitGroup) {
	time.Sleep(1*time.Second)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i:=0; i<5; i++ {
		wg.Add(1)
		go worker1(&wg)
	}
	// wg.Wait会一直阻塞main函数, until wg.Add 和 wg.Done 相互抵消完毕
	wg.Wait()

}
