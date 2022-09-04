package Base

import (
	"fmt"
	"testing"
)

type testInt func(int) bool

func isOdd(i int) bool {
	return i%2 == 1
}

func isEven(i int) bool {
	return i%2 == 0
}

// 过滤函数，通过输入的函数 filterFunc 来判断是否为目标数据进行过滤
func filter(inputArr []int, filterFunc testInt) []int {
	var outputArr []int
	for _, value := range inputArr {
		if filterFunc(value) {
			outputArr = append(outputArr, value)
		}
	}
	return outputArr
}

func TestFuncArgs(t *testing.T) {
	intputArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Even number list is -> %v\n", filter(intputArr, isEven))
	fmt.Printf("Odd number list is -> %v\n", filter(intputArr, isOdd))
}
