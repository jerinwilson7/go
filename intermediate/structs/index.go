package main

import (
	"fmt"
)

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell //anonymous field
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func main() {

	p1 := Person{
		firstName: "Jerin",
		lastName:  "Wilson",
		age:       78,
		address: Address{
			city:    "New York City",
			country: "United States Of America",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123456789",
			cell: "123456789",
		},
	}

	p2 := Person{
		firstName: "John",
		lastName:  "Doe",
	}

	p2.address.city = "London"
	p2.address.country = "UK"

	fmt.Println("FULL NAME", p1.fullName())

	p1.incrementAgeByOne()

	fmt.Println(p1)
	fmt.Println("P2 :", p2)
	fmt.Println(p1.cell) //we can call p1.cell no need to call p1.phoneHomeCell.cell. but for city we need to call p1.address.city not address.city
	fmt.Println(p1 == p2)

	// Anonymous struct

	user := struct {
		userName string
		email    string
	}{
		userName: "Jane",
		email:    "jane@gmail.com",
	}

	fmt.Println(user)

}
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
