package main

import "fmt"

// 一些channel的基本特性
func playChannel() {
	//// make制造的东东都是指针
	//d := make(chan bool)
	//go helloChannel(d)
	//<-d // 一直到d里面写入一些东东, 不然程序会卡在这里, 没办法向前走~
	//fmt.Println("Done :)")

	// 将数字 589分解为5, 8, 9 计算平方和, 立方和
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares, cubes)
}

// channel的存入和取出必须是连贯发生的, 不然goroutine会卡住
func helloChannel(done chan bool) {
	fmt.Println("Hello goroutine")
	done <- true
}

// 例如: 将一个数字456, 分解为: 6, 5, 4, 然后输入到dchnl chan里面
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	// close函数告诉receiver, 没有更多的值传给channel了
	//for {
	//	v, ok := <-c
	//	if ok!=false{
	//
	//	}
	//}
	// 没有close函数的话, dchnl会死锁
	//一个无缓冲的chan, 往里面不停传入值还不让死锁的话, 记得给sender写一个close函数
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}
