package main

import (
	"encoding/json"
	"fmt"
)

type vdc struct {
	Medium  string
	onearry []Hahs
	pointer *gdh
}
type Hahs struct {
	Name  string
	Class int
}

type gdh struct {
	Name  string
	class int
}

func main() {
	Total := vdc{Medium: "english", onearry: []Hahs{Hahs{Name: "sai", Class: 56}},
		pointer: &gdh{Name: "pavan", class: 56}}
	fmt.Println(Total)
	j, _ := json.Marshal(Total)
	fmt.Println(j)
	// -------------------------------->
	ms := vdc{Medium: "english", onearry: []Hahs{Hahs{Name: "sai", Class: 56}},
		pointer: &gdh{Name: "pavan", class: 56}}
	fmt.Println("ms value--------------->", ms)
	var trace map[string]interface{}//map creation
	fmt.Println("myMap value ---------->", trace)
	var mp map[vdc]interface//inserting in the new var
	data, err := json.Marshal(ms)
	if data != nil {
		fmt.Println("error")
	}
	// json.Unmarshal(data, &myMap)
	fmt.Println(trace["Medium"])
	fmt.Println(trace["Name"])
}
