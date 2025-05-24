package main

import "fmt"

var middleName = "LaRusso"

func main() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastName := "Smith"

	middleName := "Johnny"

	fmt.Println(middleName)
	// Default values

	// Numeric Types: 0
	// Boolean Types: false
	// String Types: ""
	// Pointers,slices,maps,functions, and structs: nil

	// Scope
	printName()
}

func printName() {
	firstName := "Danial"
	fmt.Println(firstName)
}
