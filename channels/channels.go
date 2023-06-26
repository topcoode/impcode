package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// channel := make(chan int)

	// fmt.Println(channel)
	// channel <- 5
	// channel <- 56
	// fmt.Println(string(<-channel))
	fix()
	execute()
}

// fatal error: all goroutines are asleep - deadlock!
func goroutines() {

	fmt.Println("frist function........")
}
func goroutinesTwo() {
	fmt.Println("second function.......")
}
func fix() {

	go goroutines()
	go goroutinesTwo()
	time.Sleep(time.Microsecond)
}

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go runner1(wg)
	go runner2(wg)

	wg.Wait()
}
