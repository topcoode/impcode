package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating:")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("type of a thanks for rating %t", input)
}
