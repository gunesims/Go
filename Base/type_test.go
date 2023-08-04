package Base

import (
	"fmt"
	"testing"
)

type NewInt int     // 将NewInt定义为int类型
type IntAlias = int // 将int取一个别名叫IntAlias

func TestType(t *testing.T) {

	// 将a声明为NewInt类型
	var a1 NewInt
	fmt.Printf("a1 type: %T\n", a1)

	// 将a2声明为IntAlias类型
	var a2 IntAlias
	fmt.Printf("a2 type: %T\n", a2)

	//	输出
	//	a1 type: Base.NewInt
	//	a2 type: int
}
