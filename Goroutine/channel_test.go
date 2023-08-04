package Goroutine

import "testing"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		//runtime.Gosched()
		println(a, v)
		total += v
	}
	c <- total
}

func TestChannel(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int, 3) // 类似于python队列，设置一个队列的长度，当队列满员时报错
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	go sum(a, c)

	x, y := <-c, <-c

	println(x, y, x+y)
}
