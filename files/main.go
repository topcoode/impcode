package main

import (
	"log"
	"os"
)

// func main() {
// 	// creating file
// 	facebook := "Great app while searching the tech senerio" //creating content in file
// 	file, err := os.Create(".\filenew.go")                   //file creation
// 	if err != nil {                                          // error statement
// 		panic(err) // panic  = shutdown the program and shows the error
// 	}
// 	length, err := io.WriteString(file, facebook) //writing the content in new file
// 	if err != nil {
// 		panic(err)
// 	} // err statement
// 	fmt.Println(length)
// 	defer file.Close() //close statement

// }
// Golang program to illustrate how to create
// an empty file in the default directory
func main() {

	myfile, e := os.Create("GeeksforGeeks.txt")
	if e != nil {
		log.Fatal(e)
	}
	log.Println(myfile)
	myfile.Close()
}
