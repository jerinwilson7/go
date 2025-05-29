package main

import (
	"fmt"
	"slices"
)

// slices are a more powerful and flexible abstraction over arrays

// Slices are reference types. Changing a slice element affects the underlying array and other slices that reference it.

// Slices do not own the data they reference. The data is part of the underlying array.

func main() {

	// we declare slices similar to array but the slices does no have fixed length
	// var sliceName[]elementType

	// var num []int
	// var num1 = []int{1, 2, 3}

	// num2 := []int{3, 4, 5}

	// slice := make([]int, 5) // Initiating a slice with capacity of 5

	a := [5]int{7, 8, 9, 4, 5}

	saSlice := a[1:3] //8,9

	fmt.Println(saSlice)

	numberArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numberArraySlice := numberArray[5:7]

	fmt.Println(numberArraySlice)
	numberArraySlice = append(numberArraySlice, 4, 456)
	fmt.Println("Slice", numberArraySlice)

	numberArraySliceCopy := make([]int, len(numberArraySlice))
	copy(numberArraySliceCopy, numberArraySlice)

	fmt.Println("SliceCopy", numberArraySliceCopy)

	for index, value := range numberArraySlice {
		fmt.Println(index, value)
	}

	// numberArraySlice[3] = 502

	// fmt.Println(numberArraySlice)

	if slices.Equal(numberArraySlice, numberArraySliceCopy) {
		println("Slice one equals to slice copy")
	} else {
		fmt.Println("Slice one is not equals to slice copy")
	}

	twoD := make([][]int, 3)

	for i := 0; i < len(twoD); i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	// slice[low:high]

	slice1 := numberArraySlice[2:4]

	fmt.Println(slice1)

	fmt.Println("CAPACITY OF THE SLICE IS ", cap(slice1))
	fmt.Println("LENGTH OF THE SLICE IS :", len(slice1))
}
