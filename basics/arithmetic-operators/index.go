package main

import "fmt"

func main() {
	var a, b int = 10, 3

	var result int

	result = a + b

	fmt.Println("ADDITION :", result)

	result = a * b

	fmt.Println("MULTIPLICATION :", result)

	result = a - b
	fmt.Println("SUBTRACTION :", result)

	result = a / b
	fmt.Println("DIVISION :", result)

	result = a % b
	fmt.Println("REMAINDER :", result)

	const p float64 = 22 / 7.0

	fmt.Println(p)

}
