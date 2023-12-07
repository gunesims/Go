package Goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(6 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 2
	}()

	select {
	case <-ch1:
		fmt.Println("Received from ch1")
	case <-ch2:
		fmt.Println("Received from ch2")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
