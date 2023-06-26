package main

import (
	"fmt"
	"net/url"
)

// const myurl string
func tinku3() {
	creatingURL := &url.URL{
		Scheme:  "https",
		Host:    "sai",
		Path:    "java",
		RawPath: "user=sai",
	}
	fmt.Println(creatingURL)
}
