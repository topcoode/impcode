// Golang program to show how to use Field
// Tags in the Definition of Struct Type
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`         // field tag for Name
	Aadhaar int    `json:"aadhaar"`      // field tag for Aadhaar
	Street  string `json:"street"`       // field tag for Street
	HouseNo int    `json:"house_number"` // field tag for HouseNO
}

func main() {

	var p Person

	p.Name = "ABCD"
	p.Aadhaar = 1234123412341234
	p.Street = "XYZ"
	p.HouseNo = 10

	fmt.Println(p)

	// returns []byte which is p in JSON form.
	jsonStr, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(jsonStr))

	// Sample JSON data
	var str = `{
		"name" : "PQRX",
		"aadhaar" : 1234123412341234,
		"street" : "XYZW",
		"house_number" : 10
	}`

	var p2 Person

	// retains values of fields from JSON string
	err = json.Unmarshal([]byte(str), &p2)
	// and stores it into p2
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(p2)
}
