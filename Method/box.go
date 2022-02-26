package Method

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

/* 颜色 */
type Color byte

// 不能使用:= 的赋值方式，:=的方式只能用在函数内部，函数外部使用则会无法编译通过
// 一般使用var方式定义全局变量
var strings = []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}

func (c Color) String() string {
	return strings[c]
}

/* 箱子*/
type Box struct {
	Width, Height, Depth float64
	Color                Color
}

// 计算体积
func (b Box) Volume() float64 {
	return b.Width * b.Height * b.Depth
}

// 此处的传入参数为指针类型，如果是实参，则传入的为原始数据的拷贝，达不到改变数据的目的
func (b *Box) Setcolor(c Color) {
	b.Color = c
	//*b.Color = c		// 与上一行代码效果一样
}

/* 箱子列表 */
type BoxList []Box

func (bl BoxList) BiggestColor() Color {
	volume := 0.0
	color := Color(WHITE)
	for _, box := range bl {
		if bv := box.Volume(); bv > volume {
			volume = bv
			color = box.Color
		}

	}
	return color
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].Setcolor(BLACK)
		//(&bl[i]).Setcolor(BLACK)		// 与上一行效果一样，GO自动转换指针变量
	}
}

func ExecuteBox() {
	boxes := BoxList{
		Box{Width: 4, Height: 4, Depth: 4, Color: RED},
		Box{Width: 10, Height: 10, Depth: 1, Color: YELLOW},
		Box{Width: 1, Height: 1, Depth: 20, Color: BLACK},
		Box{Width: 10, Height: 10, Depth: 1, Color: BLUE},
		Box{Width: 10, Height: 30, Depth: 1, Color: WHITE},
		Box{Width: 20, Height: 20, Depth: 20, Color: YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].Color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("\nLet's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].Color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

func TestBox() {
	fmt.Println("Box.Test() begin!")
	ExecuteBox()
	fmt.Println()
}
