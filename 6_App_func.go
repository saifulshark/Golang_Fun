package main

import "fmt"

//SOLID, here S: Single Responsibility Principle
//Our today's session is out there on the topic of Single Responsibility Principle

func main() {
	greetings()
	name := getName()
	num1, num2 := getNum()
	sum := sum(num1, num2)
	output(name, sum)
	goodByesComplement()
}

func greetings() {
	fmt.Println("Welcome to the Application.")
}
func getName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}
func getNum() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter num1: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter num2: ")
	fmt.Scanln(&num2)
	return num1, num2
}
func sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func output(name string, sum int) {
	fmt.Println("Congratulations! Sir, ", name, ".")
	fmt.Println("Summation of your numbers is: ", sum)
}
func goodByesComplement() {
	fmt.Println("Thanks for using the Application!See you again...")
	fmt.Println("Goodbye.")
}
