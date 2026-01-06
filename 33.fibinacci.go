package main

import "fmt"

// func fibonacci(n int) int {
// 	f0 := 0
// 	f1 := 1
// 	fn := f0 + f1
// 	for i := 2; i <= n-2; i++ {
// 		fn = fn + (i - 1)
// 	}
// 	return fn
// }
// func init() {
// 	fmt.Println("Fibonacci Server is Initializing... Get ready to boom!")
// }
func main() {
	// var n int
	// fmt.Println("Enter your index to see it's fibonacci seris sum number!")
	// fmt.Scanln(&n)
	// if n == 0 {
	// 	fmt.Println(0)
	// } else if n == 1 {
	// 	fmt.Println(1)
	// } else {
	// 	fmt.Println(fibonacci(n))
	// }
	var n int
	fmt.Println("Enter your index to see it's fibonacci seris sum number!")
	fmt.Scanln(&n)
	a, b := 0, 1
	fmt.Println("Fibonacci Series:...")
	for i := 0; i <= n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}

}
