package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	//MakeGetReq()
	//EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	Courses := []course{
		{"ReactJs Bootcamp", 299, "Princedev.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 399, "Princedev.in", "bcd123", []string{"full-dev", "js"}},
		{"Angular Bootcamp", 299, "princedev.in", "hit123", nil},
	}
	finalJson, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename":"ReactJs Bootcamp",
		"Price":299,
		"website":"agarwal8789@gmail.com",
        "tags":["web-dev","js"]
	}`)
	var Courses course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &Courses)
		fmt.Printf("%#v\n", Courses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}
}
func MakeGetReq() {
	// make GET request
	response, error := http.Get("https://reqres.in/api/products")
	if error != nil {
		fmt.Println(error)
	}

	// print response
	fmt.Println(response)
}
