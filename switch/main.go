package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for random no
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(5) + 1
	fmt.Println("", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("pokko")
	case 2:
		fmt.Println("dorro")
	case 3:
		fmt.Println("spider")
	case 4:
		fmt.Println("dragon")
	case 5:
		fmt.Println("badman")
	case 6:
		fmt.Println("suman")
	default:
		fmt.Println("error")
	}

	// 4
	// dragon
	// 2
	// dorro

}
