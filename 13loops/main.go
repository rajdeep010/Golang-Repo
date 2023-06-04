package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// it returns the index
	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days{
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	fmt.Println("")
	for _, day := range days{
		fmt.Printf("index value is %v\n", day)
	}

	temp := 1
	for temp < 10 {

		if temp == 2 {
			goto lco
		}

		// will cause infinite loop in go
		// if temp == 5{
		// 	continue
		// }

		// works perfectly
		if temp == 5{
			temp++;
			continue
		}

		if temp == 8{
			break
		}

		fmt.Println("Value is : ", temp)
		temp++
	}


	lco:
		fmt.Println("Jumping at https://codioshare.web.app")
}