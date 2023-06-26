package main

// methods
// structs to methods
import "fmt"

type stds struct {
	name       string
	class      int
	fathername string
}

func pre() {

	sai := stds{"sai", 55, "vevaik"}
	fmt.Println(sai)
	sai.fish() //            ::::main Name = differnt func
}
func (u stds) fish() { //imp
	fmt.Println(u.fathername) //            :::::methods =  variables

}

// // Go program to illustrate the
// // method with struct type receiver
// package main

// import "fmt"

// // Author structure
// type author struct {
// 	name	 string
// 	branch string
// 	particles int
// 	salary int
// }

// // Method with a receiver
// // of author type
// func (a author) show() {

// 	fmt.Println("Author's Name: ", a.name)
// 	fmt.Println("Branch Name: ", a.branch)
// 	fmt.Println("Published articles: ", a.particles)
// 	fmt.Println("Salary: ", a.salary)
// }

// // Main function
// func main() {

// 	// Initializing the values
// 	// of the author structure
// 	res := author{
// 		name:	 "Sona",
// 		branch: "CSE",
// 		particles: 203,
// 		salary: 34000,
// 	}

// 	// Calling the method
// 	res.show()
// }
