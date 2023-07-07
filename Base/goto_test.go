package Base

import "testing"

func TestGoto(t *testing.T) {
	i := 0
here:
	println(i)
	i++
	if i < 16 {
		goto here
	}
}
