package main

import (
	"fmt"
	"time"
)

func f(from string) { // these is written because to implement how many times in main func
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") //------------------1

	go f("goroutine") //------------------2

	go func(msg string) { //-------------------3 go func(){....()}()
		fmt.Println("the value of msg :", msg)
	}("going")

	time.Sleep(time.Second) //wait group
	fmt.Println("done")
}
