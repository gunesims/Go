package Base

import (
	"fmt"
	"os"
)

func SumAndProduct1(A, B int) (int, int) {
	add := A + B
	Multiplied := A * B
	return add, Multiplied
}

// go中的函数能够在设置返回值时给返回的参数命名
func SumAndProduct2(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func argFunc(args ...int) {
	for index, n := range args {
		fmt.Printf("args[%v] = %v\n", index, n)
	}
}

func pointFunc(a *int) int {
	*a = *a + 1
	return *a
}

// 使用defer，在函数结束后逆序执行
func deferFunc() bool {
	f, err := os.Open("./src/Base/file")
	defer f.Close()

	defer fmt.Println("This is first defer.")

	if err != nil {
		panic(err)
	}

	defer fmt.Println("This is second defer.")

	b1 := make([]byte, 10)
	line, err := f.Read(b1)
	fmt.Printf("%d bytes: %s \n", line, string((b1)))

	defer fmt.Println("This is third defer.")
	return true
}

func funcTest() {
	//x := 3
	//y := 4
	//
	//xPLUSy, xTIMESy := SumAndProduct1(x, y)
	//fmt.Println("With SumAndProduct1：")
	//fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	//fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
	//
	//fmt.Println()
	//xPLUSy, xTIMESy = SumAndProduct2(x, y)
	//fmt.Println("With SumAndProduct2：")
	//fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	//fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
	//
	//fmt.Println()
	//argFunc(1, 2, 2, 3, 4, 5, 67, 8)
	//
	//fmt.Println()
	//b := 3
	//fmt.Println("Before pointFunc, b -> ", b)
	//b1 := pointFunc(&b)
	//fmt.Printf("After pointFunc, b -> %v, b1 -> %v ", b, b1)

	deferFunc()
}
