package main

import (
	"fmt"
	"maps"
)

// A map maps keys to values

func main() {

	// var mapVariable map[key]ValueType

	// creating map using make
	// mapVariable = make(map[keyType]variableType)

	// initializing map using map literal
	// mapVariable = map[ketType]ValueType {
	// key1 = value1
	//key2 = value2
	// }

	myMap := map[string]int{
		"BANK": 25100,
		"TAX":  81000,
	}

	fmt.Println(myMap)

	myMapMake := make(map[string]int)

	fmt.Println(myMapMake)

	myMapMake["NIJU"] = 56420
	myMapMake["JOHN"] = 12450

	// delete(myMapMake, "NIJU")

	// clear(myMapMake)

	value, unknownValue := myMapMake["JOHN"] // this return the value if it exits  and  flag which is true if the value exits or false if no value exits with the given key. this is because if we give invalid key the value returns 0 but we still don't know whether the value exits or not so we can use the flag for safe

	fmt.Println(value, unknownValue)
	fmt.Println(myMapMake)

	//if we want to check the whether a value exits for the give key without value we do the following:

	_, okFlag := myMapMake["JOHN"]
	fmt.Println(okFlag)

	myMap2 := map[string]int{"student1": 43, "student2": 89}

	fmt.Println(myMap2)

	if maps.Equal(myMap2, myMapMake) {
		fmt.Println("THE GIVEN MAPS ARE IDENTICAL")
	} else {
		fmt.Println("THE GIVEN MAPS ARE DIFFERENT")
	}

	// for i, j := range myMap {
	// 	fmt.Println(j, i)
	// }

	for key, value := range myMap {
		fmt.Printf("your %s balance is %d \n", key, value)
	}

	_, ok := myMap["BANK"]

	if ok {
		fmt.Println("THE KEY EXISTS IN THE PROVIDE MAP")
	} else {
		fmt.Println("KEY DOESNT EXISTS ON THE MAP")
	}

	fmt.Println(len(myMap2))

	nestedMap := make(map[string]map[string]int)

	nestedMap["JOHN"] = myMap

	fmt.Println(nestedMap)
}
