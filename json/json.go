package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s1 := &Student{"Amanda", 12}
	bytes_s1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes_s1)) //{"Name":"Amanda","Age":12}
	var s2 string = `{"name":"Judy","age":13}`
	var stu = &Student{}
	fmt.Println("the value is stu:----------->", string(s2))
	err = json.Unmarshal([]byte(s2), stu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal: Name: %s, Age: %d\n", stu.Name, stu.Age) //Unmarshal: Name: Judy, Age: 13
}
