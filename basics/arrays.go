package basics

import "fmt"

func main() {

	// var arrayName [size]elementType
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 1
	numbers[1] = 22
	numbers[2] = 333
	numbers[3] = 4444
	numbers[4] = 55555
	fmt.Println(numbers)

	fruits := [4]string{"apple", "chery", "orange", "banana"}
	fmt.Println("Fruits array: ", fruits)

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index ", i, " = ", numbers[i])
	}

	for index, value := range fruits {
		fmt.Printf("Index %v, Value: %v\n", index, value)
	}

	a, b := someFunction()

	fmt.Printf("A: %v, B: %v\n", a, b)
	fmt.Println()

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()

}

func someFunction() (int, int) {
	return 1, 2
}
