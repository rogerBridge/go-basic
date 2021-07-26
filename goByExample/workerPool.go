package main

import (
	"fmt"
	"time"
)

// worker函数将jobs里面的value拿出来, 将处理后的结果发送给results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {

	// 为了使用 worker 工作池并且收集其的结果，我们需要 2 个通道。
	const numJobs = 20
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// 这个启动了2个worker, 从jobs channel里面取数值
	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}
	// 收集results channel里面的所有value, 这个也确保了所有的job已经被处理完毕
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
