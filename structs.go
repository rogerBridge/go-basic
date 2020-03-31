package main

import "fmt"

// struct 名称大写, 可以从别的包之中搜索到
// struct field大写, 可以从别的包之中拿到这个field
// 同一个包之内, 可以小写, 如果需要在不同的包之间传递struct和struct field, 那就老实大写吧

type Student struct {
	firstname string
	lastname string
	age string
	height string
	advantage string
}

func outStudentInfo(s *Student)  {
	fmt.Println("Before", s)
	s.advantage = "Computer Science"
	fmt.Println("After", s)
}

func inStudentInfo() *Student {
	s1 := Student{
		firstname: "An",
		lastname:  "Ruifeng",
		age:       "24",
		height:    "180",
		advantage: "Sing, Jump, Rap",
	}
	return &s1
}
