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
	
	
	fmt.Println("")
	rajdeep.GetStatus()

	fmt.Println("")
	rajdeep.NewMail()

	// it was the same that we passed so the copy was created
	fmt.Printf("Name is %v and email is %v.", rajdeep.Name, rajdeep.Email)
}

// note: first letter is capital
// so we can export it, else not
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

	// oneAge int	
	// as its first letter not capital
	// we can't export it
}

// method is nothing but function when we call it in a class (class it not there in go)
// here we can make a function outside the struct
// we just have to pass a struct

// now it can be said as a method
func (u User) GetStatus() { 
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("User of this user is changed to :", u.Email)
}