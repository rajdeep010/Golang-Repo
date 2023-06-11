package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)



func main() {
	fmt.Println("Handling get Request")
	// PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "Rajdeep")
	data.Add("lastname", "Mallick")
	data.Add("email", "rajdeepmallick999@gmail.com")

	res, err := http.PostForm(myurl, data)
	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	

	fmt.Println("Content is :", string(content))
}