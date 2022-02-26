package Interface

import (
	"fmt"
	"sort"
)

type Sequence []int

// -----------------------------------第一种方式---------------------------------
//func (s Sequence) Len() int {
//	return len(s)
//}
//
//func (s Sequence) Less(i, j int) bool {
//	return s[i] > s[j]
//}
//
//func (s Sequence) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}

//func (s Sequence) String() string {
//	sort.Sort(s) // 此处传入的是interface，需要实现Len()/Less(i, j int)/Swap(i, j int)三个接口
//	str := "["
//	for i, elem := range s {
//		if i > 0 {
//			str += " "
//		}
//		str += fmt.Sprint(elem)
//	}
//	return str + "]"
//}

// -----------------------------------第二种方式，强制转换---------------------------------
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func SortTest() {
	s := Sequence{1, 23, 14, 65, 96, 7, 78}
	fmt.Println("After sort:", s)
}
