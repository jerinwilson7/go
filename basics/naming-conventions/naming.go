package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase
	//   Structs, Interfaces,Enums

	//Eg. CalculateArea, UserInfo, NewHTTPRequest

	// snake_case

	// For naming variables, constants, or file names

	// UPPERCASE

	// for naming CONSTANTS

	//camelCase

	// used for naming variables or identifiers

	// MAINTAIN  CONSISTENCY ACROSS THE PROJECT IN NAMING VARIABLES, CONSTANTS, IDENTIFIERS, NAMING FILES ETC

	const MAX_RETRIES = 5

	var employeeID = 1001

	fmt.Println("EMPLOYEE ID ::", employeeID)
}
