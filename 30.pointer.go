// package main

// import "fmt"

// //pointer exercising by varible
//
//	func main() {
//		num := 40
//		address := &num
//		fmt.Println(*address + num)
//		*address = 99
//		fmt.Println(address)
//		fmt.Println(num)
//	}
package main

import "fmt"

func main() {
	var a int
	a = 5
	fmt.Println(a)
	p := &a
	*p = 54
	fmt.Println(a)
}
