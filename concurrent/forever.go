package main

import (
	"log"
	"time"
)

func main() {
	forever := make(chan bool)
	go func() {
		//forever <- true
		for {
			log.Println("Hello")
			time.Sleep(time.Second)
		}
	}()
	<-forever
}
