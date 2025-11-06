package main

import "fmt"

type Category struct {
	BookName    string
	Description string
	Price       float64
}

func main() {
	var book1 Category
	book1 = Category{ // Initialization- These are the instance or object of struct
		BookName:    "Go Programming",
		Description: "A book about Go programming language",
		Price:       29.99,
	}
	fmt.Println(book1)
	fmt.Println(book1.BookName)
	fmt.Println(book1.Description)
	fmt.Println(book1.Price)
	book2 := Category{
		BookName:    "Python Programming",
		Description: "A book about Python programming language",
		Price:       39.99,
	}
	fmt.Println(book2)
	fmt.Println(book2.BookName)
	fmt.Println(book2.Description)
	fmt.Println(book2.Price)
}
