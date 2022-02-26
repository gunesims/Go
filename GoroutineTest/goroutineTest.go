package GoroutineTest

import "runtime"

func sayHi(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		println(s)
	}
}

func GoroutineTest() {
	go sayHi("Goroutine say hi.")
	sayHi("Sample say hi.")
	//println()
}
