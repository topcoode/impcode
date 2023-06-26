// Go program to illustrate how
// to close a channel using for
// range loop and close function
package main

import "fmt"

func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks"
	}
	close(mychnl)
}

func main() {

	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		fmt.Println("the value of res value is ", res)
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}

//dead lock----------->
