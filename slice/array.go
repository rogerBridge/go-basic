package main

import (
	"fmt"
)

func main() {
	// 证明切片是引用类型的方法
	a := make([]string, 2)
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
}

func declareAndUse() {
	var a [100]int
	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
}

func reverse(s []string) {
	var intel string
	intel = s[1]
	s[1] = s[0]
	s[0] = intel
}
