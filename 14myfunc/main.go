package main

import "fmt"

func main() {
	fmt.Println("Learning functions in go")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is : ", result)

	proResult, msg := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("My msg is : ",msg)
}

//         type of arguments  return type
func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	// for i := range values{
	// 	total += values[i]
	// }

	for _, val := range values {
		total += val
	}

	return total, "Hi, pro result function"
}

func greeter() {
	fmt.Println("Namaste from golang")
}
