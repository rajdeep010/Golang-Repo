package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handling get Request")
	PerformGetRequest()
}

func PerformGetRequest() {
	
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code : ", res.StatusCode)
	fmt.Println("Content length", res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(content))

	// We can use the "strings" package 
	// that creates a builder for string
	// just another way of making a string and work with it,
	
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is :", byteCount)
	fmt.Println("Response is: ", responseString.String())
}
