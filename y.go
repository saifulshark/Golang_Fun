package main

import (
	"fmt"

	supercalculator "z.com/superCalculator"
)

func main() {
	// Calculate and print results from individual functions
	fmt.Println("Addition:", supercalculator.Add(10, 5))
	fmt.Println("Subtraction:", supercalculator.Subtract(10, 5))
	fmt.Println("Multiplication:", supercalculator.Multiply(10, 5))
	fmt.Println("Division:", supercalculator.Divide(10, 5))

	// // You can also use the package variables
	// fmt.Println("\nUsing package variables:")
	// fmt.Printf("Number1: %d, Number2: %d\n", supercalculator.Number1, supercalculator.Number2)

	// // Call the Supercalculator function (though it currently doesn't return values)
	// fmt.Println("\nCalling Supercalculator function:")
	// supercalculator.Supercalculator()
}
