package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var collegeName string = "xius" //variable
	type reader struct {
		class1 string
		class2 string
	}

	input := bufio.NewReader(os.Stdin)
	fmt.Println("enter the Class:")
	reader, _ := reader.ReadString('\n')
	if reader <= 1 {
		fmt.Println("class1")
	} else if reader >= 2 {
		fmt.Println("class2")
	} else {
		fmt.Println("errormessage")
	}
	fmt.Println("Good class", input)
	fmt.Println(collegeName)
	// structure()
	// fmt.Println("the class are :",class)
}

// func structure() {
// 	class := []string{"fristclass", "second class", "third class"}
// 	fmt.Println("the class are :", class)
// }
// if i := 0 ; i > 2 ; i++{
// 	fmt.Println("",class)
// }
