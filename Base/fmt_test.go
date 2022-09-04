package Base

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	a := 10
	fmt.Printf("测试 a = %v\n", a)
	fmt.Printf("表示为二进制: %b\n", a)
	fmt.Printf("该值对应的unicode码值: %c\n", a)
	fmt.Printf("表示为十进制: %d\n", a)
	fmt.Printf("表示为八进制: %o\n", a)
	fmt.Printf("表示为十六进制，使用a-b: %x\n", a)
	fmt.Printf("表示为十六进制，使用A-F: %X\n", a)
	fmt.Printf("表示为Unicode格式：U+1234，等价于”U+%%04X”: %U\n", a)
	fmt.Printf("%q\n", a)

	println("==========================")

	b := 12.34
	fmt.Printf("测试 b = %v\n", b)
	fmt.Printf("无小数部分、二进制指数的科学计数法，如-123456p-78: %b\n", b)
	fmt.Printf("科学计数法，如-1234.456e+78: %e\n", b)
	fmt.Printf("科学计数法，如-1234.456E+78: %E\n", b)
	fmt.Printf("有小数部分但无指数部分，如123.456: %f\n", b)
	fmt.Printf("根据实际情况采用%%e或%%f格式（以获得更简洁、准确的输出）: %g\n", b)
	fmt.Printf("根据实际情况采用%%E或%%F格式（以获得更简洁、准确的输出）: %G\n", b)

}
