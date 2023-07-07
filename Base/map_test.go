package Base

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	//var numbers map[string]int
	numbers := make(map[string]int) // make用于内建类型（map、slice 和channel）的内存分配
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println("\nnumbers[\"three\"] -> ", numbers["three"]) // 读取数据.

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	fmt.Println("Before delete, rating -> ", rating, rating)
	delete(rating, "C")
	fmt.Println("After delete, rating -> ", rating, rating)
}
