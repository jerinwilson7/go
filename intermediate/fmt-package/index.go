package main

import "fmt"

func main() {
	fmt.Print("HELLO")
	fmt.Print("WORLD")
	fmt.Print(1234)

	fmt.Println("HELLO")
	fmt.Println("WORLD")
	fmt.Println(123)

	name := "John"
	age := 78

	fmt.Printf("The age of %s is %d \n", name, age)

	s := fmt.Sprint("HELLO WORLD IN GOLANG", 12345, "kl")

	fmt.Println(s)

	b := fmt.Sprintln("JERIN", "WILSON", "789456123")

	fmt.Println(b)

	// var customer string
	// var customerAge int

	fmt.Println("Enter our name and age")
	fmt.Scan(&name, &age)
	fmt.Printf("The age of customer %s is %d", name, age)
}
