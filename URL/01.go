package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	fmt.Println(resp)
	if err != nil {
		log.Fatal("An error occurs while handling url", err)
	}
	marshal, err := json.MarshalIndent("http://example.com/", "", "")
	fmt.Println(marshal)
	if err != nil {
		panic(err)
	}

}
