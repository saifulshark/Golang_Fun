// package main

// import "fmt"

// const tax = 200

// func main() {
// 	call()
// }

// func call() {
// 	value1 := outer()
// 	value1()
// 	fmt.Println("Result from closure value1:", value1())
// 	value2 := outer()
// 	value2()
// 	value2()
// 	fmt.Println("Result from closure value2:", value2())
// }

// func init() {
// 	fmt.Println("init function called first")
// }

// func outer() func() {

// 	name := "John Rockofeller: (A Riba King Player.)"
// 	age := 25
// 	money := 1000
// 	chanda := 10
// 	fmt.Println("Initial Name= %s, Age= %d, Money= %d", name, age, money)

//		updatedShow := func() {
//			additional := 500
//			money = money + additional - tax - chanda
//			fmt.Println("Updated Name= %s, Age= %d, Money= %d", name, age, money)
//		}
//		return updatedShow
//	}
package main

import "fmt"

const tax = 200

func main() {
	call()
}

func call() {
	value1 := outer()
	value1()
	fmt.Println("Calling value1 again:")
	value1()

	fmt.Println("\nCreating new closure value2:")
	value2 := outer()
	value2()
	value2()
	fmt.Println("Calling value2 one more time:")
	value2()
}

func init() {
	fmt.Println("init function called first")
}

func outer() func() {
	name := "John Rockofeller: (A Riba King Player.)"
	age := 25
	money := 1000
	chanda := 10
	fmt.Printf("Initial Name= %s, Age= %d, Money= %d\n", name, age, money)

	updatedShow := func() {
		additional := 500
		money = money + additional - tax - chanda
		fmt.Printf("Updated Name= %s, Age= %d, Money= %d\n", name, age, money)
	}
	return updatedShow
}
