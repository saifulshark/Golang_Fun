package main

import "fmt"

func checkLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	var YearC int
	fmt.Println("Enter the year to check it is Leap year?")
	fmt.Scanln(&YearC)
	if checkLeapYear(YearC) {
		fmt.Println(YearC, "is a Leap year!!!")
	} else {
		fmt.Println(YearC, "is not a leap yr.")
	}
}
