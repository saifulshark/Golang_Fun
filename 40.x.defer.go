package main

import "fmt"

func calculate() (result int) {
	result = 0
	fmt.Println("first", result) //0

	show := func() {
		result = result + 10
		fmt.Println("defer", result) //15
	}
	defer show()

	result = 5
	fmt.Println("second", result) //5
	return result                 //15
}

func calc() int {
	result := 0
	fmt.Println("first", result) //0

	show := func() {
		result = result + 10
		fmt.Println("defer", result) //15
	}
	defer show()

	result = 5
	fmt.Println("second", result) //5
	return result                 //5
}

func main() {
	fmt.Println("calculate result:", calculate())
	fmt.Println("calc result:", calc())
}
