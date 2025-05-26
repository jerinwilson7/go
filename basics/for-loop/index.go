package main

import "fmt"

func main() {
	// Simple iteration over a range:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iteration over collection
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for index, value := range numbers {
		if index == 2 {
			break
		}
		fmt.Printf("Index %d, value %d \n", index, value)
	}
	fmt.Println("CONTINUE")
}
