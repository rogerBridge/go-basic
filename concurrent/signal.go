package main

import "fmt"

func main() {
	startChan := make(chan int)
	chanA := make(chan int)
	chanB := make(chan int)
	chanC := make(chan int)
	go showA(startChan, chanA)
	go showB(chanA, chanB)
	go showC(chanB, chanC)
	for i := 0; i < 100; i++ {
		startChan <- i
		<-chanC
	}
	startChan <- STOP
	<-chanC
}

const STOP = -1

func showA(preChan chan int, nextChan chan int) {
	for v := range preChan {
		if v == STOP {
			nextChan <- STOP
			return
		}
		fmt.Println("A")
		nextChan <- v
	}
}

func showB(preChan chan int, nextChan chan int) {
	for v := range preChan {
		if v == STOP {
			nextChan <- STOP
			return
		}
		fmt.Println("B")
		nextChan <- v
	}
}

func showC(preChan chan int, nextChan chan int) {
	for v := range preChan {
		if v == STOP {
			nextChan <- STOP
			return
		}
		fmt.Println("C")
		nextChan <- v
	}
}
