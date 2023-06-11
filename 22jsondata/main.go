package main

import (
	"encoding/json"
	"fmt"
)

type course struct{

	// don't give a space after the json:
	// eg: `json: "coursename"` will show error
	// write `json:"coursename"`
	
	Name string `json:"coursename"`	// it says when we will pass data of struct type in json the key-value pair will contain the key name as "coursename"
	Price int	
	Platform string	`json:"website"`
	Password string	`json:"-"`	// we put - to not show this in json data
	Tags []string	`json:"tags,omitempty"`	// omitempty k aage v space maat dena,, and it says that if this field is empty don't show this field at all
}

func main() {
	fmt.Println("Welcome to JSON ")
	EncodeJson()
}

func EncodeJson(){
	myCourses := []course{
		{"ReactJS", 300, "interviewbit.com", "abc123", []string{"web-dev", "js"}},
		{"Angular", 400, "rajdeep.com", "xyz123", []string{"web-dev", "api"}},
		{"MERN DEV", 500, "hello.com", "abc123", []string{}},
	}

	// package this data as JSON data

	// we always have to pass an interface
	// finalJson, err := json.Marshal(myCourses)

	// for clear printing
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	// try this one if you want XD 
	// finalJson, err := json.MarshalIndent(myCourses, "rajdeep", "\t")

	if(err != nil){
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}