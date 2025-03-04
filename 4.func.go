package main

import "fmt"

func main() {
	a := 10
	b := 20
	add, sub, mul, div := calculator(a, b)
	fmt.Println("Addition of a and b is ", add)
	fmt.Println("Subtraction of a and b is ", sub)
	fmt.Println("Multiplication of a and b is ", mul)
	fmt.Println("Division of a and b is ", div)
}

func calculator(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
