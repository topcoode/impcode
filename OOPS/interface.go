// Golang program to illustrate the
// concept of interfaces
package main

import (
	"fmt"
)

// defining an interface
type Sport interface {

	// name of sport method
	sportName() string
}

// declaring a struct
type Human struct {

	// defining struct variables
	name  string
	sport string
}

// function to print book details
func (h Human) sportName() string {

	// returning a string value
	return h.name + " plays " + h.sport + "."
}

// main function
func main() {

	// declaring a struct instance
	human1 := Human{"Rahul", "chess"}

	// printing details of human1
	fmt.Println(human1.sportName())

	// declaring another struct instance
	human2 := Human{"Riya", "carrom"}

	// printing details of human2
	fmt.Println(human2.sportName())
}
