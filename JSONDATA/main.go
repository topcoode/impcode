package main

import (
	"encoding/json"
	"fmt"
)

type books struct {
	Maths    string `json:"mathematics"` //can change alias...
	English  string
	science  string
	password string   `json:"-"`               //its didnt show the password
	Tags     []string `json:"tags ,omitempty"` //didnt decode the data.
}

func sai() {
	study := []books{
		{"addition", "THELIFE", "tomatoproduction", "history", []string{"topology", "final"}},
		{"subtraction", "THEMODEL", "potatoproduction", "religion", []string{"kk", "yellow"}},
		{"multi", "THEWAR", "chilliproduction", "MODERN", []string{"jj", "pink"}},
	}
	// fmt.Print(study)
	// package this data as JSON data
	finaljson, err := json.MarshalIndent(study, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("the json data is %s", finaljson)
}

// outputt:he json data is [
/*{
		"Maths": "addition",
		"English": "THELIFE",
		"Tags": [
			"topology",
			"final"
		]
	},
	{
		"Maths": "subtraction",
		"English": "THEMODEL",
		"Tags": [
			"kk",
			"yellow"
		]
	},
	{
		"Maths": "multi",
		"English": "THEWAR",
		"Tags": [
			"jj",
			"pink"
		]
	}
]*/
