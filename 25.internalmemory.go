package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Sum is:", z)
	}
	add(5, 6)
	add(p, a)
}
func main() {
	a := 20
	call()
	fmt.Println(a)
}
func init() {
	fmt.Println("Hello from init function")
}
