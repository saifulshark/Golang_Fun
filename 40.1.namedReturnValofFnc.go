package main

import "fmt"

func namedReturn(a int, b int) (ret int) {
	//or can be written as ret:=a+b but here in return state only would use int
	ret = a + b
	return // implicit return of ret____or Named return value

}
func main() {
	result := namedReturn(3, 4)
	fmt.Println("Result is:", result) // Result is: 7
}
