package main

import "fmt"

func main() {
	// function didn't call directly
	// we have to have to give variable to function name
	// & give the numbers.
	fmt.Println("______________________________________")
	// ............................
	// college()
	// ...........................
	add := adder(22, 55, 55)
	fmt.Println(add)
	// .................................
	multi := multiplication(55, 55)
	fmt.Println(multi)
	// ................................
	stringvalue := stringjust("printing")
	fmt.Println(stringvalue)
	// ....................
	intstringvalue, a := intstringaddition(55, "add")
	fmt.Println(intstringvalue, a)

	// ..................................................
	proresult, proresult2, proresult3 := proadder(2, 5, 5, 5, 0)
	fmt.Println("", proresult)
	fmt.Println("", proresult2)
	fmt.Println("", proresult3)

}
func college1() {

	fmt.Println("hello namastha")
}

// .........................................................................................................
// function signature
// two values
func adder(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

// hello namastha
// 77
// ...............................................................................
func multiplication(num int, num2 int) int {
	return num * num
	/*______________________________________
	hello namastha
	132
	3025
	*/
}

// .....................................................
// string functions
func stringjust(stringA string) string {
	return stringA
}

/*
printing*/
// ........................................
// string conversion in golang
// we have to write in func method.
func intstringaddition(intvalue int, stringvalue string) (int, string) {
	return intvalue, stringvalue
}

// ..................................................
func proadder(values ...int) (int, string, float32) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "hello", 56.88
}

//  17
// .....................................................
