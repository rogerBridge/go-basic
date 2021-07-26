package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	luckyList := make([]int, 5)
	for i:=0; i<5; i++ {
		luckyList[i] = rand.Intn(10)
	}
	fmt.Println(luckyList)
}
