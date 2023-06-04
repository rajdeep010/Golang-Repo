package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	// we can't take the { to next line
	// must be in the same line LOL

	// even we can't take else if or else on the next line
	if loginCount < 10 {
		result = "Regular user"
	}else if loginCount > 10{
		result = "Something else"
	}else {
		result = "Exact 10"
	}

	fmt.Println(result)

	fmt.Println("")
	if 9%2 == 0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	fmt.Println("")
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less")
	}

}