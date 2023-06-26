package main

import (
	"fmt"
	"log"
)

func main() {
	var thing string = "varchar"
	for i := 0; i < 5; i++ {
		// fmt.Println(thing)
	}

	// switch {
	// case thing == "varchar":
	// 	fmt.Println("0001")
	// case thing == "thing":
	// 	fmt.Println("micky")
	// default:
	// 	fmt.Println("invalid statement")

	// }
	for key, thing := range string(thing) {
		log.Print(key, thing)
	}
	thing()
}
func (a yellow) thing() {
	a.name = "name"
	a.class = 56
	fmt.Println(a)
	// return name,class

}

type yellow struct {
	name  string
	class int
}
