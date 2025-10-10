package main

import "fmt"

func main() {
	add(5, 10) // 5 and 10 are arguments
}

func add(x int, y int) { //x and y are parameters
	fmt.Println("Sum:", x+y)
	// return x + y
}
