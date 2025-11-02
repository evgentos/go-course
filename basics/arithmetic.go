package basics

import (
	"fmt"
)

func main() {
	// Variables eclaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substruction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const P float64 = 22 / 7.0
	fmt.Println(P)

	var m float64 = 1.0e-323
	fmt.Println(m)
	m = m / 10.0
	fmt.Println(m)
}
