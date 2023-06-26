package main

import (
	"errors"
	"fmt"
)

func tinku() {
	err := errors.New("Sample Error")
	if err != nil {
		fmt.Print(err)
	}
}
