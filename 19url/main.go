package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=reactjs&paymentid=akcny8aeba42fjjb"

func main() {
	fmt.Println("Welcome to url handling URLs")

	fmt.Println(myurl)


	// parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)	// o/p : https
	// fmt.Println(result.Host)	// lco.dev:3000
	// fmt.Println(result.Path)	// /learn
	// fmt.Println(result.Port())	// 3000
	fmt.Println(result.RawQuery)	// course=reactjs&paymentid=akcny8aeba42fjjb

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)	// url.Values  (kind of key values)

	fmt.Println(qparams["course"])	// [reactjs]

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	// how to make a url from its parts
	// don't forget the &
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=rajdeep",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}