package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// tell size and type of content in it
	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	// fruits[2] = "lichi"
	fruits[3] = "mango"

	fmt.Println("Fruit list : ", fruits)
	fmt.Println("Fruit list : ", len(fruits)) 	// 4 (although filling 3 items)


	var veg = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list : ", veg)
	fmt.Println("Veg list : ", len(veg))


}