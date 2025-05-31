package main

import "fmt"

func main() {
	process()

}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVER ", r)
		}
	}()

	fmt.Println("STARTING PROCESS")
	panic("SOMETHING WENT WRONG")
}
