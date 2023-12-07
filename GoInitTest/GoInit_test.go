package GoInitTest

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init2 function --->")
}

func init() {
	fmt.Println("init1 function --->")
}

var _ int64 = setValue()

func setValue() int64 {
	println("function setValue() ---->")
	return 1
}

func TestGoInit(t *testing.T) {
	fmt.Println("main ---->")
}
