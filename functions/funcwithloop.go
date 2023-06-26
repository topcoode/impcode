package main

import "fmt"

func man() {
	returntype /* returntype2, returntype3*/ := college(55, 88, 6)
	fmt.Println(returntype)
	// fmt.Println(returntype2)
	// fmt.Println(returntype3)
	for i := 0; i <= 10; i++ {
		fmt.Println(returntype)
	}
}
func college(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

/*149
149
149
149
149
149
149
149
149
149
149
149
*/
