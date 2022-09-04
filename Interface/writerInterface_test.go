package Interface

import (
	"fmt"
	"testing"
)

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p

	l := len(slice)

	if l+len(data) > cap(slice) {
		newSlice := make(ByteSlice, (l+len(data))*2) // 给切片设置大小
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}

	*p = slice
	return len(data), nil
}

func TestWriterInterface(t *testing.T) {
	var b ByteSlice
	fmt.Fprintf(&b, "Hello %v", "Tom")
}
