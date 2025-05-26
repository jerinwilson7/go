package main

import "fmt"

func main() {
	i := 1

	for i <= 5 {
		fmt.Println("Iteration ", i)
		i++
	}

	sum := 0

	// FOR AS WHILE WITH BREAK STATEMENT
	for {
		sum += 10
		fmt.Println("SUM ITERATION :", sum)
		if sum >= 50 {
			break
		}
	}

	// FOR AS WHILE WITH CONTINUE STATEMENT

	num := 0
	for num <= 100 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("ODD NUMBER : ", num)
		num++
	}
}
