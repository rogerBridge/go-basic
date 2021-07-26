package main

import (
	"fmt"
)

// 0,1    1,2,3,5,8,13...
func Fibonacci(i int) int {
	if i <= 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

func FibonacciMine(flag int) int {
	initS := []int{0, 1}
	for i:=0; i<flag; i++ {
		initS = append(initS, initS[i]+initS[i+1])
	}
	fmt.Println(initS[2:len(initS)-1])
	return initS[len(initS)-1]
}

func main() {
	//fmt.Println(FibonacciMine(1000))
	fmt.Println(Fibonacci(1000))
}