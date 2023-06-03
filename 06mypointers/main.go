package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int	// nil (initially)
	// fmt.Println("value of ptr is : ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is:", myNumber)
}