package main

import (
	"fmt"
) // Add this import

var aaa int8 = 127
var xxx int64 = 9223372036854775807
var ppp uint8 = 250

var flags bool = false
var line string = "I Love NYC"
var nums float64 = 13.1221111

func main() {

	rune := '❤'

	fmt.Println(aaa) // Capital P in Println
	fmt.Println(xxx)
	fmt.Println(ppp)
	fmt.Printf("%c\n", rune)
	fmt.Println(rune)
	fmt.Printf("%T\n", aaa)
	fmt.Printf("%v", flags)
	fmt.Printf("%s", line)
	fmt.Printf("%.4f\n", nums)
}
