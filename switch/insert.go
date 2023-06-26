package main import (
	    "encoding/json"
	    "fmt"
	) type MyStruct1 struct {
	    Field1 string
	    Field2 int
	} type MyStruct2 struct {
	    Field3 string
	} type MyStruct struct {
	    Field4 string
	    Field5 int
	    Arr    []MyStruct1
	    Ptr    *MyStruct2
	} func main() {
	    // create a structure variable
	    s := MyStruct{
	        Field4: "value1",
	        Field5: 123,
	        Arr: []MyStruct1{
	            MyStruct1{Field1: "value2", Field2: 456},
	            MyStruct1{Field1: "value3", Field2: 789},
	        },
	        Ptr: &MyStruct2{Field3: "value4"},
	    }     // convert the structure to a map
	    m := make(map[string]interface{})
	    m["Field4"] = s.Field4
	    m["Field5"] = s.Field5
	    arr := make([]interface{}, len(s.Arr))
	    for i, v := range s.Arr {
	        arr[i] = v
	    }
	    m["Arr"] = arr
	    if s.Ptr != nil {
	        ptr := make(map[string]interface{})
	        ptr["Field3"] = s.Ptr.Field3
	        m["Ptr"] = ptr
	    }     // convert the map to JSON
	    j, err := json.Marshal(m)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	    fmt.Println(string(j))
	}