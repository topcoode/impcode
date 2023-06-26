package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://github.com/microsoft/vscode-go/issues/2858"

func saivevek() {
	http.Get(url)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the type of a url %T", response)
	// response.Body.Close()
	// the type of a url *http.Response
	datatype, err := ioutil.ReadAll(response.Body)
	// func ioutil.ReadAll(r io.Reader) ([]byte, error)
	if err != nil {
		panic(err)
	}
	content := string(datatype)
	log.Println("content", content)
}

// http packages
/*resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})*/
