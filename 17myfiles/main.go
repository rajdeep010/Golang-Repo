// we can only work with text files
// csv and pdf need spcl libraries

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Leaning about files in golang")

	content := "This needs to go in a file - rajdeep.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	// panic shut down the program and show error
	// 	panic(err)
	// }

	checkNilErr(err)

	len, err := io.WriteString(file, content)

	// if err != nil{
	// 	panic(err)
	// }

	checkNilErr(err)


	fmt.Println("Length is : ", len)
	defer file.Close()


	readFile("./mylcogofile.txt")
}


func readFile(filename string) {
	// creation of a file is "os" operation
	// so using os package

	// but for more manipulation we have io package
	// and we read data in byte format
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil{
	// 	panic(err)
	// }

	checkNilErr(err)

	fmt.Println("Text data inside the file is : ", databyte)
	fmt.Println("Text data inside the file is : ", string(databyte))

}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}