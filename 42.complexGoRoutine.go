package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main function started")
	go fmt.Println("hello this is Islam Saiful-5")
	// time.Sleep(5 * time.Second)
	goRountine()
	time.Sleep(5 * time.Second)
	fmt.Println("main function ended")
}
func goRountine() {
	go fmt.Println("hello this is saiful")
	fmt.Println("hello world")
	go fmt.Println("hello this is saiful2")
	go fmt.Println("hello this is saiful3")
	go fmt.Println("hello this is saiful4")
	fmt.Println("bye world")
	// time.Sleep(5 * time.Second)
}

// Output: main function started,hello world,bye world
