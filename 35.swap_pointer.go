package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter a=")
	fmt.Scanln(&a)
	fmt.Println("Enter b=")
	fmt.Scanln(&b)

	fmt.Println("Before swap: a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)
}

func swap(a *int, b *int) {
	*a = *a + *b // a now contains sum of both
	*b = *a - *b // b = (a+b) - b = original a
	*a = *a - *b // a = (a+b) - original a = original b
}
