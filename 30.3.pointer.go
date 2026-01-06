// pointer using function
package main

import "fmt"

func updateValue(val *int) {
	*val = 100
}

func main() {
	num := 42
	fmt.Println("Before:", num)
	updateValue(&num)
	fmt.Println("After:", num)
}
