package main

import (
	"log"
	"time"
)

func main() {
	for ;; {
		go sendMsg()
		t0 := time.NewTimer(1*time.Minute)
		<-t0.C
	}
}

//
func sendMsg() {
	log.Println("sending...")
	time.Sleep(10*time.Second)
	log.Println("sent")
}