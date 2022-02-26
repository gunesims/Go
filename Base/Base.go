package Base

import "fmt"

func execute() {
	arrTest()
	mapTest()
	ifTest()
	gotoTest()
	forTest()
	switchTest()
	funcTest()
	funcArgTest()
	structTest()
}

func Test() {
	fmt.Printf("Base.Test() begin!\n\n")
	execute()
	fmt.Println()
}
