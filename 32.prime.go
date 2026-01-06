// package main

// import "fmt"

// func main() {
// 	var numberInt int
// 	fmt.Println("Enter the number to check int weather Prime or Not!")
// 	fmt.Scanln(&numberInt)

// 	prime := true
// 	if numberInt == 1 {
// 		prime = false
// 	} else if numberInt == 2 || numberInt == 3 || numberInt == 5 || numberInt == 7 {
// 		prime = true
// 	} else {
// 		for i := 2; i < 8; i++ {
// 			if numberInt%i == 0 {
// 				prime = false
// 				break
// 			}
// 		}
// 	}

//		if prime {
//			fmt.Println(numberInt, "is a Prime number!")
//		} else {
//			fmt.Println(numberInt, "is NOT a Prime number...Zzz")
//		}
//	}
package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Print("Enter a number (0 to n numbers): ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		prime := true
		if n <= 1 {
			prime = false
		} else {
			for i := 2; i*i <= n; i++ { // efficient version
				if n%i == 0 {
					prime = false
					break
				}
			}
		}

		if prime {
			fmt.Println(n, "is a Prime number")
		} else {
			fmt.Println(n, "is NOT a Prime number")
		}
	}
}
