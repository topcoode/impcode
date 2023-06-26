package main

import "fmt"

func main() {
	// declaring maps using var keyword
	college := map[string]string{}
	college["name"] = "vdc"
	college["address"] = "xyz"
	fmt.Printf("printing the key and values:%v\n", college)
	// initializing th map using the make() func
	student := make(map[string]int)
	fmt.Printf("empty string:%v\n", student)
	student["name"] = 56
	student["rollno"] = 56
	fmt.Printf("using make keyword in maps:%v\n", student)
	college2()
	studentdata := student["name"]
	fmt.Printf("studentdata value is :%v\n", studentdata)
	college3()
	loopswithmaps()
	looptesting()
	length()
}
func college2() {

	student := map[int]string{ //interface{} must be there
		56: "fifty six",
		52: "fifty two",
		53: "fifty three",
	}
	// fmt.Println(student)
	//accessing the values in maps in golang
	namecarry := (56)
	fmt.Printf("accessing the values:%v\n", namecarry)
	//changing the vlue
	student[56] = "changevalue"
	fmt.Printf("changed value:%v\n", student)
	//delete the the map
	delete(student, 53)
	fmt.Printf("after deleting the values are:%v\n", student)
}
func college3() {
	//existing are not in maps
	collegename := map[string]interface{}{
		"name":          "sai",
		"addressname":   "poiiii",
		"addressnumber": 56,
	}
	fmt.Println(collegename)
	value, exist := collegename["monkey"]
	if exist {
		fmt.Printf("existed the value :%v %t \n ", value, exist)
	} else {
		fmt.Printf(" the existed values are :%t \n", exist)
	}
}
func loopswithmaps() {
	loops := map[string]string{
		"name:":  "sai",
		"class:": "second",
		"bench:": "second",
	}
	// fmt.Printf("the loops start with:------------>%v\n", loops)
	for key, value := range loops {

		fmt.Printf("loops in maps: %v %v \n", key, value)
	}

}
func looptesting() {
	loops := map[string]string{
		"name:":  "sai",
		"class:": "second",
		"bench:": "second",
	}
	// fmt.Printf("the loops start with:------------>%v\n", loops)
	for value := range loops {

		fmt.Printf("loops with values: %v \n", value)
	}

}
func length() {

	length := map[int]string{

		1: "sai",
		2: "vinay",
	}
	fmt.Printf("the length of the map :%v", len(length))
}
