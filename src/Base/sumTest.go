package Base

import "fmt"

// go中的函数能够在设置返回值时给返回的参数命名
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func sumTest() {
	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
