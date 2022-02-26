package main

import (
	"GoTest/Base"
	"GoTest/Fibonacci"
	"GoTest/GoroutineTest"
	"GoTest/Interface"
	"GoTest/WebTest"
)

func main() {
	Base.Test()

	Fibonacci.Test()

	Interface.WriterInterfaceTest()
	Interface.SortTest()
	Interface.Test()

	GoroutineTest.GoroutineTest()
	GoroutineTest.ChanTest()

	WebTest.WebServerStart()
}
