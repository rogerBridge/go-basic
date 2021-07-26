package main

import "fmt"

func main() {
	msg := make(chan string, 2)
	msg <- "hello"
	msg <- "world"
	close(msg) // 如果不使用close的话, range会继续请求msg, msg没有新的消息注入, 会导致死锁
	for item := range msg {
		fmt.Println(item)
	}
}
