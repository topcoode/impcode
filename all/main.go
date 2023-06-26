package main

// // // // import "fmt"

// // // // func main() {
// // // // 	var a string = "name"
// // // // 	for i := 0; i >= 7; i++ {
// // // // 		fmt.Println("",a[i])

// // // // 	}

// // // // }

// // // package main

// // // import "fmt"

// // // func main() {
// // // 	var a int32 = 99
// // // 	for i := 0; i <= 7; i++ {
// // // 		fmt.Println("", a)
// // // 	}

// // // // }
// // package main

// // import "fmt"

// //	func main() {
// //		a := 20.5
// //		b := 20.5
// //		if a == b {
// //			fmt.Println("the value of int and float value is")
// //		} else if a != b {
// //			fmt.Println("the value of int is")
// //		} else {
// //			fmt.Println("the value of float is")
// //		}
// //	}
// package main

// import "fmt"

// func main() {
// 	months := []string{"jan", "feb", "mar"}
// 	for _, redirect := range months {
// 		fmt.Println(redirect)
// 	}

// }
// import (
// 	"runtime"
// 	"strconv"
// )

// func MessageWithFileLine(message string, skip int) string {
// 	_, file, line, ok := runtime.Caller(1 + skip)
// 	if ok {
// 		message += " " + file + ":" + strconv.Itoa(line)
// 	}
// 	return message
// }
import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) { //handler function
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) { //handler function
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func stringooooo() {

	// http.HandleFunc()//getRoot(), getHello()
	// // ungetRooti, getHello, err := http.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//--------------- to handle the function
	//   call the(http.ListenAndServe) function    ///////////For Runingserver
	//   to start the (server) and listen for requests.

}
