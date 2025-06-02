package main

import "fmt"

func main() {
	// Create a closure (function with state) using adder()
	// when adder gets initialized it prints the initial value of i and stores the return func in sequence variable
	sequence := adder()

	// Call the closure multiple times
	// Each call increases the internal counter `i` by 1
	fmt.Println(sequence()) // Output: 1
	fmt.Println(sequence()) // Output: 2
	fmt.Println(sequence()) // Output: 3
	fmt.Println(sequence()) // Output: 4
	fmt.Println(sequence()) // Output: 5

	// Create a new, independent closure from adder()
	sequence2 := adder()

	// This starts from 0 again, because it's a new instance
	fmt.Println(sequence2()) // Output: 1
	fmt.Println(sequence2()) // Output: 2
	fmt.Println(sequence2()) // Output: 3
	fmt.Println(sequence2()) // Output: 4

	subtractor := func() func(int) int {
		countdown := 99
		return func(i int) int {
			countdown -= i
			return countdown
		}
	}()
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(18))
}

// adder returns a closure function that increments and returns a counter
func adder() func() int {
	i := 0 // Local variable captured by the returned closure; initialized only when adder() is called

	// Just to show the initial value when adder() is called
	fmt.Println("Previous value of i is", i)

	// Return a closure that captures `i` and increments it each time it is called
	return func() int {
		i++ // Increment the captured value
		fmt.Println("Add the value of i by one")
		return i // Return the updated value
	}

}
