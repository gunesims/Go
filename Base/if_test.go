package Base

import (
	"fmt"
	"testing"
)

func scanfx() int {

	var x int
	fmt.Scanf("%d", &x)

	return x
}

func TestIf(t *testing.T) {
	//var x int
	//fmt.Scanf("%d", &x)
	//
	//if x > 10 {
	//	fmt.Println("x is greater than 10")
	//} else {
	//	fmt.Println("x is less than 10")
	//}

	// 使用if中条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if y := scanfx(); y > 10 && y < 20 {
		fmt.Println("10 < y < 20")
	} else if y > 20 && y < 30 {
		fmt.Println("20 < y < 30")
	} else {
		fmt.Println("y > 30")
	}

}
