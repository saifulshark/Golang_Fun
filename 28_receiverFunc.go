// package main

// import "fmt"

// // Category struct definition
// type Category struct {
// 	BookName    string
// 	Description string
// 	Price       float64
// }

// // Regular function to print book details
// func printFunctionBook1(book Category) {
// 	fmt.Println("Regular function output:")
// 	fmt.Println(book)
// 	fmt.Println("Book Name:", book.BookName)
// 	fmt.Println("Description:", book.Description)
// 	fmt.Println("Price:", book.Price)
// 	fmt.Println()
// }

// // Receiver function that MODIFIES the struct using pointer
// // It takes a new price as parameter and updates the book's price
// func (bk Category) printFunctionBook2(newPrice float64, bkk string) {
// 	fmt.Println("Before modification:")
// 	fmt.Println("Book Name:", bk.BookName)
// 	fmt.Println("Old Price:", bk.Price)
// 	fmt.Println()

// 	// Modify the price not using the pointer
// 	bk.Price = newPrice
// 	bk.BookName = bkk
// 	fmt.Println("After modification (not using pointer receiver):")
// 	fmt.Println("Book Name:", bk.BookName)
// 	fmt.Println("New Price:", bk.Price)
// 	fmt.Println("Memory address:", &bk)
// 	fmt.Println()
// }

// // Receiver function that takes a Category parameter and updates ALL fields
// func (bk *Category) updateBook(newBook Category) {
// 	fmt.Println("Before update:")
// 	fmt.Printf("Book: %+v\n", *bk)
// 	fmt.Println()

// 	// Update all fields using pointer
// 	bk.BookName = newBook.BookName
// 	bk.Description = newBook.Description
// 	bk.Price = newBook.Price

// 	fmt.Println("After update (pointer modified the original):")
// 	fmt.Printf("Book: %+v\n", *bk)
// 	fmt.Println("Memory address:", bk)
// 	fmt.Println()
// }

// func main() {
// 	// Create first book
// 	book1 := Category{
// 		BookName:    "Go Programming",
// 		Description: "A book about Go programming language",
// 		Price:       29.99,
// 	}

// 	// Create second book
// 	book2 := Category{
// 		BookName:    "Python Programming",
// 		Description: "A book about Python programming language",
// 		Price:       39.99,
// 	}

// 	fmt.Println("=== Example 1: Regular Function (Pass by Value) ===")
// 	printFunctionBook1(book1)

// 	fmt.Println("=== Example 2: Receiver Function with Parameter ===")
// 	// Pass new price to receiver function - it will modify book2
// 	book2.printFunctionBook2(99.99, "Python Orelly Crash Course")

// 	// Print book2 again to see the change persisted
// 	fmt.Println("=== Example 2.1 ===")
// 	fmt.Println("Checking if book2 was modified:")
// 	fmt.Printf("Book2 Price is now: %.2f\n\n", book2.Price)
// 	fmt.Println(book2)

// 	fmt.Println("=== Example 3: Update Entire Struct ===")
// 	// Create a new book with different values
// 	newBookData := Category{
// 		BookName:    "Rust Programming",
// 		Description: "A book about Rust programming language",
// 		Price:       49.99,
// 	}

// 	// Update book1 with new data using pointer receiver
// 	book1.updateBook(newBookData)

//		// Print book1 to verify changes persisted
//		fmt.Println("Checking if book1 was completely updated:")
//		fmt.Printf("Book1: %+v\n", book1)
//	}
package main

import "fmt"

type Category struct {
	BookName string
	Price    float64
}

// VALUE receiver - does NOT modify original
func (bk Category) changeWithValue(newPrice float64) {
	bk.Price = newPrice // only changes the COPY
	fmt.Println("Inside value receiver:", bk.Price)
}

// POINTER receiver - DOES modify original
//inshort:- &=address and *=value
//&ampersand= for address(i.e. #1004aeiou memory location addess) and *asterisk=value for pointing to address for storing value or allocation.
func (bk *Category) changeWithPointer(newPrice float64) {
	bk.Price = newPrice // changes the ORIGINAL
	fmt.Println("Inside pointer receiver:", bk.Price)
}

func main() {
	book := Category{
		BookName: "Go Programming",
		Price:    29.99,
	}

	fmt.Println("Original price:", book.Price)
	fmt.Println()

	// Test 1: Value receiver
	fmt.Println("--- Using Value Receiver ---")
	book.changeWithValue(99.99)
	fmt.Println("After value receiver:", book.Price) // Still 29.99!
	fmt.Println()

	// Test 2: Pointer receiver
	fmt.Println("--- Using Pointer Receiver ---")
	book.changeWithPointer(99.99)
	fmt.Println("After pointer receiver:", book.Price) // Changed to 99.99!
}
