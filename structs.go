package main

import "fmt"

// struct 名称大写, 可以从别的包之中搜索到
// struct field大写, 可以从别的包之中拿到这个field
// 同一个包之内, 可以小写, 如果需要在不同的包之间传递struct和struct field, 那就老实大写吧

type Student struct {
	firstname string
	lastname  string
	age       string
	height    string
	advantage string
}

func outStudentInfo(s *Student) {
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

// 两个struct可以比较相等性吗?
// 当field不可比较时, 两个struct不相等, 当包含的field可比较&&比较结果相等时, 两个struct相等
func isStructEqual() {
	stu1 := Student{
		age: "24",
	}
	stu2 := Student{
		firstname: "",
		lastname:  "",
		age:       "24",
		height:    "",
		advantage: "",
	}
	switch {
	case stu1 == stu2:
		{
			fmt.Println("equal")
		}
	default:
		fmt.Println("not equal")
	}
}
