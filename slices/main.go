package main

import (
	"fmt"
	"reflect"
)

func main() {

	// RemoveIndex()
	college := []string{"sai", "frnd", "is"}
	fmt.Println(college)
	fmt.Println(reflect.ValueOf(college).Kind())
	//slices with the loops
	for key, value := range college {

		fmt.Println("the value are by using the range are : ", key, value)
	}
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("the values are :", i)
	// }
	//--------------------------------->
	var intSlice = make([]int, 10)
	var strSlice = make([]string, 10, 20)
	fmt.Printf("intSlice Len: %v Cap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
	// ------------------------------->
	fmt.Print("from here the func starts-------------------------------->")
	//declare slice using "new" keyword
	var new = new([50]string)[0:10]
	fmt.Println(new)
	fmt.Println(reflect.ValueOf(new).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	// adding the values to the slices
	var addingthevalues = make([]int, 5, 5)
	addingthevalues[0] = 56
	addingthevalues[1] = 54
	addingthevalues[2] = 56

	fmt.Println("adding values are :", addingthevalues)
	//changing the item value
	addingthevalues[0] = 86
	addingthevalues[1] = 85
	addingthevalues[2] = 96
	fmt.Println("changing the values", addingthevalues)
	delete()
	coping()
}

// deleting the values
func delete() {
	var strSlice = []string{"a", "b", "c", "d", "e"}
	defer fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	defer fmt.Println(strSlice)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// coping the function
func coping() {
	var copya = []string{"apple", "ball", "cat"}
	var copyb = make([]string, 2, 2)
	copyb[0] = "dog"
	copyb[1] = "elephant"
	fmt.Println(copya, copyb)
	copy(copyb, copya)
	fmt.Println(copyb)
}
