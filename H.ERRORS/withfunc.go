package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil {
		if errors.Is(err, filepath.ErrBadPattern) {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}
