package main

import "fmt"

func main() {
	message := "HELLO WORLD!"

	for index, value := range message {
		// fmt.Println(index, value)

		fmt.Printf("index: %d value: %c \n", index, value)
	}
}
