package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	lang := make(map[string]string)

	lang["JS"] = "Javascript"
	lang["PY"] = "python"
	lang["RB"] = "ruby"
	lang["CP"] = "C++"

	fmt.Println("List of all languages: ", lang)
	fmt.Println("JS short for: ", lang["JS"])

	// delete ruby from map
	delete(lang, "RB")
	fmt.Println("Now the list is: ", lang)

	// loops are intersting in golang
	// %v is for value
	for key, value := range lang {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	fmt.Println("")
	// ignore anything with _ sign in golang if we are not sure about it
	for _, value := range lang {
		fmt.Printf("For value is %v\n", value)
	}
}