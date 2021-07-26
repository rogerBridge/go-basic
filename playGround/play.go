package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10*time.Second)
	//time.Sleep(5*time.Second)
	go func() {
		for i:=0;;i++{
			fmt.Println(i)
		}
	}()
	<-timer.C
}