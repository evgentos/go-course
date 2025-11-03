package basics

import "fmt"

func main() {

	fruit := "apple"

	switch fruit {
	case "orange":
		fmt.Println("It's a orange")
		fmt.Println("---------")
	case "apple":
		fmt.Println("It's an apple")
	case "cucumber", "potato":
		fmt.Println("It's a vegetable")
	default:
		fmt.Println("Unknown fruit")

	}

}
