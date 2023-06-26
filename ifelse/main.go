package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	incomingcalls := 77
	var result string
	if incomingcalls <= 10 {
		result = "HO no!!!!!!!!!"
	} else if incomingcalls > 70 {
		result = " STOP"
	} else {
		result = "!!!!"
	}
	fmt.Println(result)
	// ------------------------------------------------------------
	// taking input and checking ifelse statement
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating:")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	var x int
	if x == 90 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}
