package main

import (
	"log"
	"os"
)

func main() {
	// 把原始数据进行加密
	content := os.Args[1:]
	if len(content) != 2 {
		log.Fatalln("输入的参数长度错误!")
	}
	cmd := content[0]
	data := content[1]
	if cmd == "--encrypt" {
		encrypt(data)
	}else if cmd == "--decrypt" {
		decrypt(data)
	}else {
		log.Fatalln("参数错误, 请输入: --encrypt 或者 --decrypt")
	}
}
