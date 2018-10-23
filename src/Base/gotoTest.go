package Base

func gotoTest() {
	i := 0
here:
	println(i)
	i++
	if i < 16 {
		goto here
	}
}
