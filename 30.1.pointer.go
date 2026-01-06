// pointer using array
package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	var ptr [3]*int
	for i := 0; i < len(arr); i++ {
		ptr[i] = &arr[i]
	}
	fmt.Println(*ptr[0], *ptr[1], *ptr[2])
}
