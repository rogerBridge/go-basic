package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func playBufferdChannel() {
	deadLock()
}

// Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.
// 缓冲channel当capacity满了, push blocking, 缓冲channel当capacity空了, pop blocking

// 死锁, channel pop了, 没有push, 死锁, 或者channel push了, 没有pop, 也会死锁
// 违背push-pop一致性, 就会死锁(buffered channel当buffer为空时, 和普通channel一样)

func deadLock() {
	n := make(chan string, 2)
	n <- "Hello"
	n <- "World"
	n <- "!" // channel没有pop元素, 没有空间往里面push了, deadlock
	fmt.Println("Dead lock!")
}

// capacity 和 length
func capacityAndLength() {
	c := make(chan string, 3)
	c <- "Hello"
	c <- "World"
	fmt.Println(cap(c), len(c)) // 3, 2
	<-c
	fmt.Println(len(c))
}

// Wait Group
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Start goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //Done方法: counter=counter-1
}

func playWaitGroup() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1) // 当counter为0时, 所有的goroutine block
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished")
}

// worker poll, 多个worker一起结束
type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digit(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 需要完成的工作
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 工人每次处理一件工作
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digit(job.randomno)}
		results <- output
	}
	wg.Done() //goroutine - 1
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1) // goroutine + 1
		go worker(&wg) // close(results)
	}
	wg.Wait()
	close(results)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func playWorkerPool() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)

	go result(done)

	noOfWorkers := 20
	createWorkerPool(noOfWorkers)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
