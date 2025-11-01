package main

import "fmt"

func main() {
	add(5, 10) // 5 and 10 are arguments(values are arguments=> a comes first for arguments)
}

func add(x int, y int) { //x and y are parameters(paramenters pass values from arguments)
	fmt.Println("Sum:", x+y)
	// return x + y
}
