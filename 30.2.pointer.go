// pointer exercising with struct
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"Alice", 30}
	ptr := &p
	fmt.Println("Name:", p.name)
	fmt.Println("Name:", ptr.name)
	fmt.Println("Age:", (*ptr).age)
	fmt.Println(*ptr)
}
