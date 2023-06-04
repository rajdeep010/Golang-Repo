package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO Web request")

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of %T", response)
	// Response is of *http.Response

	// so basically we are getting a reference i.e. we can change it further

	fmt.Println("")
	// fmt.Println("The response is : ", response)


	defer response.Body.Close()
	// caller's responsibility to close the connection


	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}

	content := string(databytes)

	// prints the HTML content of the web page
	fmt.Println("Content is: ", content)
}