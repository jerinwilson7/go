package main

import "fmt"

// As soon as we encounter a panic noting after the panic will be executed

func main() {
	//panic(interface{})
	process(-1)
	process(10)
}

func process(input int) {
	if input < 0 {
		panic("THE INPUT MUST BE A NON-NEGATIVE NUMBER")
	} else {
		fmt.Println(input)
	}
}
