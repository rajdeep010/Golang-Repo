package main

import (
	"encoding/json"
	"fmt"
)

type course struct {

	// don't give a space after the json:
	// eg: `json: "coursename"` will show error
	// write `json:"coursename"`

	Name     string `json:"coursename"` // it says when we will pass data of struct type in json the key-value pair will contain the key name as "coursename"
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // we put - to not show this in json data
	Tags     []string `json:"tags,omitempty"` // omitempty k aage v space maat dena,, and it says that if this field is empty don't show this field at all
}

func main() {
	fmt.Println("Welcome to JSON ")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
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

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS",
			"Price": 300,
			"website": "interviewbit.com",
			"tags": [
					"web-dev",
					"js"
			]
		}
	`)

	var mycourse course

	checkValid := json.Valid(jsonDataFromWeb)
	
	if(checkValid){
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	}else{
		fmt.Println("Json was not valid")
	}

	// some cases where you juyt want to add data to key value

	// we know the key will be a string but the value we don't know 
	// it can be anything
	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
