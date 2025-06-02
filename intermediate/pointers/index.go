package main

import "fmt"

// pointers holds the memory address of a value
// pointers are the variables that store the memory address of another variable

func main() {

	// syntax: var ptr *int

	a := 10
	ptr := &a

	*ptr = *ptr + 10
	fmt.Println(a)

	i := 45
	l := 1200

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &l

	*p = *p / 10
	fmt.Println("THE VALUE OF L IS ", l)

	fmt.Println(*p) //dereferencing a pointer

	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// 	return
	// }

	modifyValue(p)

	fmt.Println("AFTER ::", i)

}

func modifyValue(ptr *int) {
	*ptr++
}
