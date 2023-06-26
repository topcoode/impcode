package main

import "fmt"

func main() {
	//pointers
	var name string = "xius"
	var p *string
	p = &name
	fmt.Println(name)
	fmt.Println(p)
	fmt.Println(*p)

	var A int
	var z *int = &A
	fmt.Println(z)
	var month string = "feb"
	var adress *string = &month
	fmt.Println(adress)
	fmt.Println(*adress)

}
