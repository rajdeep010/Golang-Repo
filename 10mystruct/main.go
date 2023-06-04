package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, no super or parent
	rajdeep := User{"Rajdeep", "rajdeep@go.dev", true, 22}

	fmt.Println(rajdeep)

	// get a detailed view of struct
	fmt.Printf("Rajdeep's detials are: %+v\n", rajdeep)
	fmt.Println("")
	fmt.Printf("Name is %v and email is %v.", rajdeep.Name, rajdeep.Email)
}

// note: first letter is capital
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
