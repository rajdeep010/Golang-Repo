package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// if we write this way without giving the size
	// then we have to initialize on the go
	var fruits = []string{}
	fmt.Printf("Type of slices %T", fruits)


	var arr = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("The array is : ", arr)

	// slices have append method
	// also increases the size of slice
	arr = append(arr, "Mango", "Banana")
	fmt.Println("The array is : ", arr)

	// kind of substring
	arr = append(arr[1:])
	fmt.Println("The arrya is :", arr)

	// last is not inclusive
	arr = append(arr[1:3])
	fmt.Println("The arrya is :", arr)

	// start from 0
	// and before 3
	arr = append(arr[:3])
	fmt.Println("The arrya is :", arr)


	// make works on memory management
	// makes a 4 size slice of integer
	// as we are making it of size 4
	// when access index 4 error shown
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777	// cause error

	// as soon as we use append
	// it reallocates the memory
	// and all value are accomodated
	highScores = append(highScores, 555, 444, 222)

	fmt.Println(highScores)

	// sort function in slices and not in array
	// sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted((highScores)))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted((highScores)))


	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	// kind of slice(as javascript) from index 2 and append the index+1 to all
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}