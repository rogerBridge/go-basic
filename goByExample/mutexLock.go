package main
import (
	"fmt"
	"sync"
)
var x  = 0
// 对全局变量x进行 +1 的操作

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	// 确保同一时刻, 只有一个goroutine可以访问到x
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done() // 报告给waitGroup
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
