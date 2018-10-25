package Method

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 一般函数的定义是
// func area(r Rectangle) float64 {
//		return r.width * r.height
// }
// 下面这种方式实现的method，是Rectangle结构专属的函数
// method area() 是依赖于某个形状(比如说Rectangle)来发生作用的。
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func ExecuteArea() {
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 40}

	c1 := Circle{10}
	c2 := Circle{20}

	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())
}

func TestArea() {
	fmt.Println("\nArea.Test() begin!")
	ExecuteArea()
	fmt.Println()
}
