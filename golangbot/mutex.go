package main

import (
	"fmt"
	"sync"
)

// 来到了传说中的互斥锁环节
// 竞争状态

//func playNoMutex() {
//	var w sync.WaitGroup
//	for i:=1; i<101; i++{
//		w.Add(1)
//		go increment(&w)
//	}
//	w.Wait()
//	fmt.Println("The result is:", x)
//}
//
//var x = 0
//
//func increment(wg *sync.WaitGroup) {
//	x += 1
//	wg.Done()
//}

func playMutex() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 1; i < 101; i++ {
		wg.Add(1)
		go increment(wg, m)
	}
	wg.Wait()
	fmt.Println("increment with mutex :)", x)
}

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}
