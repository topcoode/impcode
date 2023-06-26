package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainnnn() {
	fmt.Println("welcome to HAHS school")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the class:")
	input, _ := reader.ReadString('\n')
	fmt.Println("bye student", input)
	if i := 0 ; i < 15 ; i++  {
	fmt.Println(class[i])
}
	structure()
}
func structure() {
	class := []string{"fristclass", "secondclass", "thirdclass"}
}

