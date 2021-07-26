package main

import "fmt"

func main() {
	flag := make(chan int, 2)
	flag <- 0
	flag <- 1
	close(flag) // 往里面不停地传入值并且还不的deadlock的话, 记得使用close()
	for v := range flag {
		fmt.Println(v)
	}
}
