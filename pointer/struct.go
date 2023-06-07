package pointer

import (
	"fmt"
)

type Book struct {
	title	string
	price	float32
}

func queryBook(book *Book) {
	fmt.Println(*book)
}

func Struct() {
	var book *Book
	book = &Book{title: "Lord of the Ring", price: 35.90}
	queryBook(book)
}