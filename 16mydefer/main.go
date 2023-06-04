package main

import "fmt"

func main() {
	// just think that this line would go
	// to just before the ending } of this func
	// (just imagine this to understand)
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	// and defers work in lifo
	// so order => Two, One, World

	fmt.Println("Hello Defer")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++{
		defer fmt.Println("Value is :", i)
	}
}