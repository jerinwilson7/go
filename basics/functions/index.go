package main

import "fmt"

func main() {

	// func <functionName>(parameterList) returnType {
	// codeBlock
	// return value
	// }

	result := sum(10, 20)
	fmt.Println(result)

	// func() {
	// 	fmt.Println("HELLO ANONYMOUS FUNCTION")
	// }()

	greet := func() {
		fmt.Println("HELLO ANONYMOUS GREET FUNCTION")
	}

	greet()

	operation := sum
	result = operation(10, 98)
	fmt.Println(result)

	response := applyOperations(10, 90, sum)

	fmt.Println(response)

}

func sum(a, b int) int {
	return a + b
}
func applyOperations(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}
