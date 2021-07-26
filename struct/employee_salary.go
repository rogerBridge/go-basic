package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Employee struct {
	name string
	salary float64
}

type myTime struct {
	t time.Time
}

func (e *Employee) giveRaise(give float64) string {
	r := fmt.Sprintf("%.2f", e.salary*(1+give/100))
	//fmt.Println("员工", e.name, "调整后的薪水为:", r, "元")
	e.salary, _ = strconv.ParseFloat(r, 64)
	return r
}

func (m *myTime) show() string {
	fmt.Println(m.t.String())
	return m.t.String()
}

func main() {
	a := &Employee{"leo", 6500}
	a.giveRaise(10)
	log.Println("调整后的薪水为: ", a.salary)
	b := &myTime{t: time.Now()}
	b.show()
}
