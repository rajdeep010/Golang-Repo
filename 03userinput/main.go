package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza :")

	// comma ok || comma error syntax
	// dont' have any try catch so we keep err in varibales

	// want to read until I find a \n
	// _ is the variable for err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,",input)
	fmt.Printf("Type of rating %T",input)	// string

	// if we are not worrying about the err
	// then we use _ for err part
	// input, err just like try,catch
	// if not worry about input then write _,err
}
