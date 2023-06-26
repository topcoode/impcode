package main

import "fmt"

type author struct {
	Firstname string
	Lastname  string
}
type Book struct {
	Publisher string
	Authors   []author
}

func main() {

	var book Book
	book.Publisher = "kdfvjbxk"
	{

		book.Authors = append(book.Authors, author{})
		book.Authors[1].Firstname = "abc"
		book.Authors[1].Lastname = "mbc"
	}

	fmt.Println(book)

}
