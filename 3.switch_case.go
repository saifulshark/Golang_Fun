package main

import "fmt"

func main() {
	age := 61
	switch {
	case age == 18 || age == 19 || age == 20:
		fmt.Println("You can Vote.")
	case age < 18:
		fmt.Println("You are too young to vote.")
	case age > 20 && age < 60:
		fmt.Println("You can vote.")
	default:
		fmt.Println("Check your age.You can not vote.")
	}
}
