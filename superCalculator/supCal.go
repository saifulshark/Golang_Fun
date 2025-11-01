package supercalculator

// Number1 - exported variable accessible outside the package
var Number1 int = 101

// Number2 - exported variable accessible outside the package
var Number2 int = 20

func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	if b != 0 {
		return a / b
	}
	return 0
}
func Supercalculator() {
	Add(Number1, Number2)
	Subtract(Number1, Number2)
	Multiply(Number1, Number2)
	Divide(Number1, Number2)
}
