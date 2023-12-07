package Goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	sum := func(a []int, c chan int) {
		total := 0
		for _, v := range a {
			//runtime.Gosched()
			//println(a, v)
			total += v
		}
		c <- total
	}

	a := []int{}
	for i := 1; i < 100000; i++ {
		a = append(a, i)
	}
	c := make(chan int, 3) // 类似于python队列，设置一个队列的长度，当队列满员时报错
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	go sum(a, c)

	x, y, z := <-c, <-c, <-c

	println(x, y, z)
}

func TestChannel2(t *testing.T) {
	c := make(chan int)
	defer close(c)

	var n = 0
	go func(n *int) {
		for {
			*n += 1
			//println("Hello World")
		}
	}(&n)

	go func(i, j int) {
		time.Sleep(3 * time.Second)
		c <- i + j
	}(3, 4)

	i := <-c
	println(i)
	println(n)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestChannel3(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
