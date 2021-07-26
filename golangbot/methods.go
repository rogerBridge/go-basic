package main

import "fmt"

func (s *Student) adv() {
	fmt.Println(s.advantage)
}

// 使用函数处理: 值结构体, 指针结构体
func dealStruct(s Student) {
	s.advantage = "modified"
}

func dealStructPoint(s *Student) {
	s.advantage = "modified"
}

// Methods with no struct || user-define type
type str string

func (ss str) show() {
	fmt.Println(ss)
}
