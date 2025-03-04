package main

import "fmt"

func main() {
	welcome()
	myCourse("saifulire")
	myCourse("24s")
	myCourse("23")
}
func welcome() {
	fmt.Println("Welcome to GoLang course, Education must be free for all")
}
func myCourse(course string) {
	fmt.Println("Welcome to the course from, ", course)
}
