package main

import (
	"fmt"

	"strconv"
)

func main() {

	name := "56"

	strvcon, err := strconv.Atoi(name)

	if err != nil {

		fmt.Println("WRONG", err)

	}

	fmt.Println(strvcon)

	main2()

	main3()

	main4()

	main5()
	main6()

}

// string tooo number

func main2() {

	name2 := "53433.777"

	strvconvv, err := strconv.ParseFloat(name2, 3)

	if err != nil {

		fmt.Println("WRONG", err)

	}

	fmt.Println(strvconvv)

}

// string tooo float

func main3() {

	name3 := "true"

	true, err := strconv.ParseBool(name3)

	if err != nil {

		fmt.Println("WRONG", err)

	}

	fmt.Println("", true)

}

// string tooo boolean

// ----------------------------------------------------------------------------------------------------

// boolean to string

func main4() {

	var b bool = true

	fmt.Printf("the value is: %v", b)

	var s string = strconv.FormatBool(true)

	fmt.Printf(" the value is %v :", s)

}

func main5() {

	b := true
	convert := strconv.FormatBool(b)
	fmt.Sprintf("%v", convert)
}
func main6() {
	var apple int32 = 125
	convert := strconv.FormatInt(int64(apple), 10)
	fmt.Println(convert)
}
