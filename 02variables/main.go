package main

import "fmt"

func main() {
	var username string = "rajdeep"
	fmt.Println(username)

	// %T gives the type of variable
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var value int = 278945
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	var smallFloat float32 = 255.454452971812
	fmt.Println(smallFloat)	// will print 255.45445
	// for float64 it will be more precise
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)	// prints 0
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type/way to declare varible
	var website = "wow.com"
	fmt.Println(website)
	fmt.Printf("The type of website is %T \n", website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)
}