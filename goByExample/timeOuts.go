package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2*time.Second)
		c1 <- "result 1"
	}()
	select {
	case <-c1:
		fmt.Println("done 1")
	case <-time.After(time.Microsecond*2000400):
		fmt.Println("timeout 1")
	}
}
