package main

import (
	"fmt"

	"example.com/mathlib" //importing mathlib package
)

//scope first : {}
//for Package scope: go mod init example.com

func main() {
	fmt.Println("Hello, World!")
	mathlib.Add(3, 4)
	fmt.Println("Goodbye, World!")
}
