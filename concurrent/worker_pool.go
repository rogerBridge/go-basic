package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	taskChan := make(chan string)
	taskList := make([]string, 100)
	for i := 0; i < 100; i++ {
		taskList[i] = fmt.Sprint(i)
	}
	//fmt.Println(taskList)

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(taskChan, wg)
	}

	for _, v := range taskList {
		taskChan <- v
	}
	close(taskChan)

	wg.Wait()
}

func worker(taskChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range taskChan {
		time.Sleep(time.Second)
		fmt.Print(v)
	}
}
