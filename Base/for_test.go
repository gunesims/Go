package Base

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	// 遍历数组
	var array = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for index := 0; index < len(array); index++ {
		fmt.Printf("%v -> %v", index, array[index])
	}

	for index := range array {
		fmt.Println(array[index])
	}

	// 遍历字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	for key, value := range rating {
		fmt.Printf("rating[%v] = %v\n", key, value)
	}

}
