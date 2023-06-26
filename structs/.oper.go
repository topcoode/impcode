// package main

// import "fmt"
// func main() {
// 	Lang := Access{"fullstack", 55, true, 55.22}
// 	fmt.Printf("the core lang are: %v", Lang.Golang)
// }

// /*  c := Car{Name: "Ferrari", Model: "GTC4",
//             Color: "Red", WeightInKg: 1920}

//     // Accessing struct fields
//     // using the dot operator
//     fmt.Println("Car Name: ", c.Name)
// 	output:Car Name:  Ferrari
// 	............................
//     fmt.Println("Car Color: ", c.Color)
//   output:Car Color:  Red
//   .............................
//     // Assigning a new value
//     // to a struct field
//     c.Color = "Black"
//     fmt.Println("Car: ", c)
// output: Car:  {Ferrari GTC4 Black 1920}*/

// type Access struct {
// 	Golang string
// 	Cpp    int
// 	Python bool
// 	Java   float32
// }

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

// Main Function
func main() {
	c := Car{Name: "Ferrari", Model: "GTC4",
		Color: "Red", WeightInKg: 1920}

	// Accessing struct fields
	// using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value
	// to a struct field
	c.Color = "Black"

	// Displaying the result
	fmt.Println("Car: ", c)
}
