package Goroutine

import (
	"runtime"
	"testing"
)

func sayHi(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		println(s)
	}
}

func TestGoroutine(t *testing.T) {
	go sayHi("Goroutine say hi.")
	sayHi("Sample say hi.")
	//println()
}
