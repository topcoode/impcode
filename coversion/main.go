package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//conversions
	reader := bufio.NewReader(os.Stdin)
	//Stdin, Stdout, and Stderr are open Files pointing to the standard input,
	//standard output, and standard error file descriptors.

	//Note that the Go runtime writes to standard error for panics and crashes;
	//closing Stderr may cause those messages to go elsewhere, perhaps to a file opened later.
	fmt.Println("enter the ratings")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks:", input)
	//func (*bufio.Reader).ReadString(delim byte) (string, error)

	numRating, err := strconv.ParseFloat(input, 64)
	//fmt.Println(numRating)
	//	func strconv.ParseFloat(s string, bitSize int) (float64, error)
	if err != nil {
		fmt.Println(err, numRating)

	}

}
