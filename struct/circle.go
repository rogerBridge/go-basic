package main
//结构体类型和非结构体类型的方法

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

// 如果*circle是一个指针, go会自动解指针
func (c *circle) perimeter() float64 {
	return math.Pi*2*c.radius
}

func (c *circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func main() {
	c1 := &circle{3.2}
	fmt.Println(c1.radius, c1.perimeter(), c1.area())
	fmt.Printf("%.4f, %.4f, %.4f", c1.radius, c1.area(), c1.perimeter())
}
