package main

import (
	"fmt"
	"time"
)

func Access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel\n") //............>

	for i := range ch {
		fmt.Println(i) //.............>
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go Access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled") //.................>
	}

	time.Sleep(3 * time.Second)
}
