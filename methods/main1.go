package main

import "fmt"

type college struct {
	BranchName  string
	ClassName   string
	StudentName string
	FatherName  string
}

func main() {
	sai := college{"mpc", "loll", "sai", "ramulu"}
	fmt.Println(sai)
	sai.sai()
}
func (r college) sai() {
	fmt.Println(r.FatherName)
	return
}
