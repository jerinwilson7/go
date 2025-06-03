package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	message := "HELLO \n GO"
	message1 := "HELLO\tGO"
	message2 := "HELLO\rGO"
	rawMessage := `HELLO \n GO`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("LENGTH OF MESSAGE ONE IS : ", len(message2))
	fmt.Println(message[0]) //ASCII

	greeting := "HELLO"
	name := "ALICE"

	fmt.Println(greeting + " " + name)

	str1 := "APPLE"
	str2 := "Orange"
	str3 := "apple"

	fmt.Println(str1 < str2)
	fmt.Println(str1 > str3)

	for index, char := range str1 + message {
		fmt.Printf("THE INDEX %d HAS THE VALUE %c \n", index, char)
		fmt.Printf("%x\n", char)

	}
	fmt.Println("RUNE COUNT: ", utf8.RuneCountInString(message+str1))

	var ch rune = 'a'
	var jch rune = '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c \n", ch)
	fmt.Printf("%c \n", jch)

	clone := strings.Clone(str1)

	fmt.Println(clone)
	fmt.Println(strings.Contains("SEAFOOD", "foo"))

	cstr := string(ch)
	fmt.Println(cstr)
}
