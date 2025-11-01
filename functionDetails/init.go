package main

import "fmt"

var a int = 10

func init() {
	fmt.Println("Initial value of a:", a)
}
func functionDetails() {
	fmt.Println("Value of a:", a)
}
func main() {
	functionDetails()
}
func init() {
	fmt.Println("Initial value of a:", a)
	a = 20
}
func init() {
	fmt.Println("updated value of a:", a)
	a = 25
}
