// Golang program to illustrate the
// concept of custom types
package main

import (
	"fmt"
)

// declaring a struct
type Book struct {

	// defining struct variables
	name   string
	author string
	pages  int
}

// function to print book details
func (book Book) print_details() {

	fmt.Printf("Book %s was written by %s.", book.name, book.author)
	fmt.Printf("\nIt contains %d pages.\n", book.pages)
}

// main function
func main() {

	// declaring a struct instance
	book1 := Book{"Monster Blood", "R.L.Stine", 131}

	// printing details of book1
	book1.print_details()

	// modifying book1 details
	book1.name = "Vampire Breath"
	book1.pages = 162

	// printing modified book1
	book1.print_details()

}
