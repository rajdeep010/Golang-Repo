package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)



func main() {
	fmt.Println("Handling get Request")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with go lang",
			"price" : 0,
			"platform" : "helloworld"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll((res.Body))
	fmt.Println(string(content))
}
