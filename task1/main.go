package main

import (
	"encoding/json"
	"fmt"
)

type college struct {
	Name  string
	class int
}
type college2 struct {
	Name1 string
}
type college3 struct {
	ROLLNO string
	Class  int
	Arr    []college //Array
	Ptr    *college2 //pointer
}

func main() {
	s := college3{
		ROLLNO: "twenty",
		Class:  123,
		Arr: []college{
			college{Name: "hahs", class: 56},
			college{Name: "brs", class: 25},
		},
		Ptr: &college2{Name1: "sai"},
	}
	fmt.Println("the value of s:----------->", s)
	// ms := s
	// convert the structure to a map
	m := make(map[string]interface{})
	fmt.Println("the value of m------------->", m)
	m["sai"] = s.ROLLNO
	m["pavan"] = s.Class
	arr := make([]interface{}, len(s.Arr))
	for i, v := range s.Arr {
		arr[i] = v
	}
	m["Arr"] = arr
	// var myMap map[string]interface{}
	// data, err := json.Marshal(ms)
	// json.Unmarshal(data, &myMap)
	// if err != nil {
	// 	fmt.Printf("the value is %v", data)
	// }
	// fmt.Println(myMap["ROLLNO"])
	// fmt.Println(myMap["Class"])
	// fmt.Println("---------------->\n", myMap)
	fmt.Println("the value of m------------->", m)

	if s.Ptr != nil {
		ptr := make(map[string]interface{})
		ptr["sai"] = s.Ptr.Name1
		m["Ptr"] = ptr
	}
	// convert the map to JSON
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))
}
