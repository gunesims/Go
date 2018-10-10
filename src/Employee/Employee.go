package Employee

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s, my phone number is %s.\n", s.name, s.phone)
}

type Employee struct {
	Human
	company string
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I working at %s. Call me on %s.\n", e.name, e.company, e.phone)
}

func Execute() {
	mark := Student{Human{"mark", 18, "13669875485"}, "CQU"}
	sam := Employee{Human{"sam", 30, "16549789321"}, "yagoo"}

	mark.SayHi()
	sam.SayHi()
}

func Test() {
	fmt.Println("Employee.Test() begin!")
	Execute()
	fmt.Print("\n")
}
