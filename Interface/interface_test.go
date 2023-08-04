package Interface

import (
	"fmt"
	"strconv"
	"testing"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %v.\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la la...", lyrics)
}

type Student struct {
	Human
	school string
	loan   float32
}

// 类似于python __repr__
func (s Student) String() string {
	return "(name:" + s.name
}

type Employee struct {
	Human
	company string
	money   float32
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %v, you can call me on %v.\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func TestInterface(t *testing.T) {
	mark := Student{Human{"Mark", 12, "13634654987"}, "CQ", 32.1}
	paul := Student{Human{"Paul", 16, "999-888-555"}, "MIT", 23.3}

	sam := Employee{Human{"Sam", 23, "444-666-333"}, "google", 1000.4}
	tom := Employee{Human{"Tom", 33, "111-666-333"}, "microsoft", 10.4}

	var menInterface Men

	menInterface = mark
	fmt.Println("This is Mark, a Student:")
	menInterface.SayHi()
	menInterface.Sing("Hello world!")
	fmt.Println()

	menInterface = sam
	fmt.Println("This is Sam, a Employee:")
	menInterface.SayHi()
	menInterface.Sing("Hello world!")
	fmt.Println()

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 4)
	x[0], x[1], x[2], x[3] = mark, paul, sam, tom

	for index, value := range x {
		value.SayHi()
		value.Sing(strconv.Itoa(index))

		// 查看数据的类型是否是指定类型，value, ok = element.(T)
		if v, ok := value.(Student); ok {
			fmt.Println("value is Student -> ", v)
		}

		switch v := value.(type) {
		case Student:
			fmt.Println("value is Student -> ", v)
		case Human:
			fmt.Println("value is Human -> ", v)
		case Employee:
			fmt.Println("value is Employee -> ", v)
		}

		//value.Sing(value.(Human.name))
		fmt.Println()
	}

}
