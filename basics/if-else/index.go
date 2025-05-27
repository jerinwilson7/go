package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	// if condition {
	// block of code
	// } else if {
	// block of code
	// } else {
	// block of code
	// }

	mark := 80

	if mark >= 90 {
		fmt.Println("Your grade is A")
	} else if mark >= 80 {
		fmt.Println("Your grade is B")
	} else if mark >= 70 {
		fmt.Println("Your grade is C")
	} else {
		fmt.Println("YOU FAILED")
	}

}
