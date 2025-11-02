package basics

import "fmt"

type GoogleEmployee struct {
	FirstName string
	LastName  string
	Age       int
}

type AppleEmployee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg: CalculateArea, UserInfo, NewHTTPRequest
	// Structs: interfaces, enums

	// snake_case
	// Eg. user_id, first_name, http_request

	// UPPERCASE
	// Используется для именования констант

	// mixedCase
	// Eg: javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
