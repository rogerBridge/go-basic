package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now() // 开始计时时刻
	//slice()
	slice2()
	t1 := time.Since(t0) // 结束计时时刻
	fmt.Println(t1)
}

func slice()  {
	s := make([]int, 1000000)
	for i:= 0; i<1000000; i++{
		s[i] = i
	}
	fmt.Println(len(s))
}

func slice2()  {
	// 提高cap可以提升很多的速度哦~
	//s := make([]int, 0, 1000000)
	s := make([]int, 0)
	//var s []int
	//s[0] = 0
	for i:=0; i<1000000; i++{
		s = append(s, i)
	}
	fmt.Println(len(s))
}