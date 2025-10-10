package main

import "fmt"

func sum_Return_two_Calculator(a int, b int) (int, int, int, int) {

	sum := a + b
	difference := a - b
	mul := a * b
	division := a / b

	return sum, difference, mul, division
}
func main() {
	x := 99
	y := 199
	summation, difference, multiplication, division := sum_Return_two_Calculator(x, y)
	fmt.Println("Sum is:", summation)
	fmt.Println("Difference is:", difference)
	fmt.Println("Multiplication is:", multiplication)
	fmt.Println("Division is:", division)
}
