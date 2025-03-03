package main

import "fmt"

func main() {
	fmt.Println("hello, world!")
	age := 18
	sex := "male"
	sex = "female"
	fmt.Println(sex)
	if age < 18 && sex == "female" {
		fmt.Println("You are young,you are not eligible to marry.But you can love him in this system. What does it mean to love him? It means you also attached with SEX!!!")
	} else if age < 18 && sex == "male" {
		fmt.Println("You are young,you are not eligible to marry.But you can love her in this system. What does it mean to love her? It means you also attached with SEX!!!")
	} else if age <= 18 && sex == "female" {
		fmt.Println("You are eligible to marry anyone(him)")
	} else if age == 20 && sex == "male" {
		fmt.Println("After 20 years old aged Man can marry any girl(18+)")
	} else if age == 18 || age > 18 {
		fmt.Println("You are eligible to marry anyone,Whoever you are!!!")
	} else {
		fmt.Println("You are too Young to marry anyone")
	}
}
