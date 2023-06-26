package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}
func PerformGetRequest() {
	const myurl = "https://github.com/gorilla/mux/blob/master/doc.go"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code", response.StatusCode)
	// fmt.Println("status code", response.ContentLength)

	// status code:"status code 200"
	// READ REPONSE BODY
	var responseString strings.Builder //create struct  //transform
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, err := responseString.Write(content) //transform
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(content))

	// byte has to transform
	fmt.Println(byteCount)
	// fmt.Println(responseString.String())

}
