package basics

import "fmt"

var middleName = "Sanich"

func main() {
	// var age int
	var name1 string = "John"
	var name2 = "Jane"

	count := 10
	lastName := "Smith"

	// Default values:
	// Numeric types: 0
	// Boolean type: false
	// String type: ""
	// Pointers, slices, maps, functions, and structs: nil

	fmt.Println(name1, name2, lastName, count, middleName, printName())
}

func printName() string {
	firstName := "Michal"
	return firstName
}
