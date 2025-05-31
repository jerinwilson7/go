package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.

func main() {

	defer process(100)
	defer fmt.Println("WORLD !!!")
	fmt.Println("HELLO")
	fmt.Println("NO THIS TIME")
}

func process(i int) {
	defer fmt.Println("DEFERRED VALUE", i)
	defer fmt.Println("FIRST DEFERRED STATEMENT EXECUTED")
	defer fmt.Println("SECOND DEFERRED STATEMENT EXECUTED")
	defer fmt.Println("THIRD DEFERRED STATEMENT EXECUTED")
	i++
	fmt.Println(i)
	fmt.Println("NORMAL STATEMENT EXECUTED")
}
