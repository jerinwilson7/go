package main

import (
	"errors"
	"fmt"
)

func main() {

	//Function with multiple return value
	// func functionName(parameter1 type1,parameter2 type2,...)(returnType1,returnType2...){
	// code block
	// return returnValue1,returnValue2.....
	// }

	fmt.Println(divide(10, 20))

	result, err := compare(20, 20)
	if err != nil {
		fmt.Println("ERROR :", err)
	} else {
		fmt.Println("RESULT ::", result)
	}

}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("unable to compare the given values")
	}
}
