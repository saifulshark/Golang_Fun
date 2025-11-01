package main

import "fmt"

var ax, b int

func input() {
	fmt.Scanln(&ax, &b)
}

func zerco() {
	fmt.Println("Sum:", suma(ax, b))
	fmt.Println("Minus:", minus(ax, b))
	fmt.Println("Mul:", mul(ax, b))
	fmt.Println("Div:", div(ax, b))
}

func suma(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func mul(x int, y int) int {
	return x * y
}

func div(x int, y int) int {
	return x / y
}
