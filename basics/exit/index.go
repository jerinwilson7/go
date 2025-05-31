package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("DEFERRED STATEMENT")
	fmt.Println("STARTING THE MAIN FUNCTION")
	os.Exit(1)
	fmt.Println("END OF THE MAIN FUNCTION")
}
