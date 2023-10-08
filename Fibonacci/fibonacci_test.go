package Fibonacci

import (
	"fmt"
	"testing"
	"time"
)

func fibs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibs(n-1) + fibs(n-2)
}

func Execute() {
	n := 40
	startTime := time.Now()
	ret := fibs(n)
	useTime := time.Since(startTime)
	fmt.Printf("fibs(%d) -> %d use time -> %v\n", n, ret, useTime)
}

func TestFibonacci(t *testing.T) {
	Execute()
}
