package Method

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) SayHi() {

	fmt.Printf("Hi, I am %s, my phone number is %s.\n", h.name, h.phone)
}

type Student struct {
	Human
	school string
}

func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s, I study at %s, my phone number is %s.\n", s.name, s.school, s.phone)
}

type Employee struct {
	Human
	company string
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I working at %s. Call me on %s.\n", e.name, e.company, e.phone)
}

func ExecuteEmployee() {
	mark := Student{Human{"mark", 18, "13669875485"}, "Microsoft"}
	sam := Employee{Human{"sam", 30, "16549789321"}, "Google"}

	mark.SayHi()
	sam.Human.SayHi()
	sam.SayHi()
}

func TestEmployee() {
	fmt.Println("\nEmployee.Test() begin!")
	ExecuteEmployee()
	fmt.Println()
}
