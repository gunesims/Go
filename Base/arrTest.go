package Base

import "fmt"

func arrTest() {
	var array = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var aSlice, bSlice, cSlice []string

	aSlice = array[:5]
	fmt.Println("aSlice = array[:5]", aSlice)

	aSlice = array[5:]
	fmt.Println("aSlice = array[5:]", aSlice)

	aSlice = array[:]
	fmt.Println("aSlice = array[:]", aSlice)

	aSlice = array[3:7]
	fmt.Println("aSlice = array[3:7]", aSlice)

	bSlice = aSlice[1:3]
	fmt.Println("bSlice = aSlice[1:3]", bSlice)

	bSlice = aSlice[:5]
	fmt.Println("bSlice = aSlice[:5]", bSlice)

	bSlice = aSlice[:]
	fmt.Println("bSlice = aSlice[:]", bSlice)

	bSlice[2] = "change"
	fmt.Println("\naSlice -> ", aSlice)
	fmt.Println("bSlice -> ", bSlice)
	fmt.Println("array -> ", array)

	cSlice = array[2:4:7]
	fmt.Println("cSlice -> ", cSlice)
}
