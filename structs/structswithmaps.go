package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type College struct {
	Name          string `json: "name"`
	LicenceNO     string `json:"licenceno"`
	StudentCount  string `json: "studentCount"`
	TeachersCount string `json:"teacherscount"`
}
type student struct {
	Name         string
	HallticketNO int
	college      []College
}

func main() {
	Info := student{
		Name: "sai", HallticketNO: 56,
		college: []College{
			College{
				Name:          "pavan",
				LicenceNO:     "two",
				StudentCount:  "five",
				TeachersCount: "fiftyfive",
			}}}
	// fmt.Println(Info)
	pikkky, err := json.Marshal(Info)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pikkky))
	nikky, err := json.MarshalIndent(Info, "", "")
	fmt.Println(string(nikky))
	var out bytes.Buffer
	json.Indent(&out, pikkky, "=", "\t")

	out.WriteTo(os.Stdout)
}
