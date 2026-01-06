package main

import "fmt"

func dfr() {
	var def float32
	fmt.Println("first: ", def) //null
	def = 0.1
	fmt.Println("second: ", def) //0
	def = def + 1
	defer fmt.Println("last: ", def) //1
	def = def + 2
	fmt.Println("thrid: ", def) //3
	def = def + 1
	defer fmt.Println("second last: ", def)
}
func main() {
	dfr()
}
