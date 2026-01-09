package main

import "fmt"

func calculate() (result int) {
	result = 0
	fmt.Println("first", result) //0

	show := func() {
		result = result + 10
		fmt.Println("defer", result) //15
	}
	defer show() //15

	result = 5
	p := func(a int) {
		fmt.Println("me", a)
	}
	defer p(result) //me:5

	defer fmt.Println(result) //5

	fmt.Println("second", result) //5

	defer fmt.Println(5) //5

	return result //15
}
func main() {
	fmt.Println("calculate result:", calculate())
}
