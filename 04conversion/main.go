package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading,",input)

	// numRating = input + 1 (can't do that for not doing conversion)

	// numRating, err := strconv.ParseFloat(input, 64)

	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println("Added 1 to your rating: ", numRating+1)
	// }
	//  printing this give me the error : strconv.ParseFloat: parsing "4\r\n": invalid syntax


	// Handling a string and making it a integer
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}


}
