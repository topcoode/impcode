package main

import (
	"fmt"
)

type data struct {
	InstanceID string
	DefaultID  string
	data2      data2
}
type data2 struct {
	Maintaince string
	Tools      string
}

func main() {
	Data := data{
		InstanceID: "one",
		DefaultID:  "two",
		data2: data2{
			Maintaince: "abc",
			Tools:      "dfv",
		},
	}
	
	data2, err := fmt.Println(Data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)
}
