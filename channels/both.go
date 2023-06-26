package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	defer close(ch)
	fmt.Println("the value of ch", ch)
	go func(ch chan int) {
		time.Sleep(time.Second)
		ch <- 1
	}(ch)

	<-ch
	fmt.Println(ch)
}
