package main

import (
	"fmt"
)

func main() {
	fristslice := []string{"orange", "apple", "banana", "mango"}
	fmt.Println(fristslice)
	secondslice := []int{55, 55, 66, 66, 88, 22}
	defer fmt.Println(secondslice)
	thirdmap := map[int]string{
		1: "bringle",
		2: "tomato",
		3: "spinach",
	}
	defer fmt.Println(thirdmap)
	for id, values := range thirdmap {
		fmt.Println(id, values)
	}
}
