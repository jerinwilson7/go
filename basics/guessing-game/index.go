package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game")
	fmt.Println("I have guessed a number between 1  and 100")
	fmt.Println("Can you guess what it is")

	var guess int

	for {
		fmt.Println("Enter you guess")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("You guess is right !!!!!!!")
			break
		} else if guess > target {
			fmt.Println("You have guessed a greater number than my guess try again")
		} else {
			fmt.Println("You have guessed a smaller number than my guess try again")
		}
	}
}
