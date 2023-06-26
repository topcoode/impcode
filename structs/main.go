package main

import "fmt"

func ranjith() {
	// no inheritance & no super or parent
	Sai := Std{"sai", 55, "graduate"}
	fmt.Println(Sai)
	// {sai 55 graduate}
	fmt.Printf("the details of sai is: %+v/n", Sai)
	//the details of sai is: {Name:sai Rollno:55 Degree:graduate}
	fmt.Printf("the name: %v and the Roll no:%v and Degree:%v", Sai.Name, Sai.Rollno, Sai.Degree)
	//  name: sai and the Roll no:55 and Degree:graduate

}

type Std struct {
	// Std S- capital, std such as class.
	//fields- as export so frist as to be capital
	Name   string
	Rollno int
	Degree string
}

// How to access fields of a struct,dot (.) operator,
