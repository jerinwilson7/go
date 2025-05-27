package main

import "fmt"

func main() {
	// general syntax ::var arrayName [size]element type

	var numArray = [5]int{1, 2, 3, 4}

	numArray[3] = 30

	fmt.Println(numArray)

	fruits := [3]string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruits)

	// for i := 0; i < len(numArray); i++ {
	// 	fmt.Println("Element at index ,", i, ":", numArray[i])
	// }

	for index, value := range numArray {
		fmt.Printf("value at index %d is %d \n", index, value)
	}

	// underscore is blank identifier used to store unused value

	a, _ := someFunction()

	fmt.Println(a)

	b := 2
	_ = b

	fmt.Println("Length of the array is ", len(numArray))

	ogArray := [3]int{1, 2, 2}
	var arrayCopy *[3]int

	arrayCopy = &ogArray
	fmt.Println("ORIGINAL ARRAY :", ogArray)
	arrayCopy[0] = 100
	fmt.Println("ORIGINAL ARRAY :", ogArray)
	fmt.Println("COPY ARRAY :", arrayCopy)
}

func someFunction() (int, int) {
	return 10, 20
}
