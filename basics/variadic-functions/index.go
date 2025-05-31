package main

import "fmt"

//A variadic function in Go is a function that accepts a variable number of arguments of the same type.
// ...Ellipsis

func main() {

	// func functionName(param1 type1,param2 type2,param3 ...type3)returnType{
	// FUNCTION BODY
	// }

	fmt.Println(sum(1, 2, 3, 4, 5, 5, 6, 4, 78))
	fmt.Println(sumWithString("Sum of 1,2,3,4,5,5,6,4,78 is", 1, 2, 3, 4, 5, 5, 6, 4, 78))

}

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total = total + value
	}

	return total
}

func sumWithString(returnString string, nums ...int) (string, int) {
	total := 0
	for _, value := range nums {
		total = total + value
	}

	return returnString, total
}
