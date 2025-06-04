package main

import "fmt"

type Person struct {
	name string
	age  int
}

func createNewPerson(name string, age int) *Person {
	p := Person{name: name, age: age}

	return &p
}

func main() {

	fmt.Println(Person{name: "John", age: 22})
	fmt.Println(Person{name: "Jane", age: 20})
	fmt.Println(Person{name: "Charlie"})
	fmt.Println(&Person{name: "Don", age: 34})
	fmt.Println(createNewPerson("Kevin", 25))

	s := Person{name: "Alice", age: 12}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 90

	fmt.Println(sp)
	fmt.Println(s)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: false,
	}

	fmt.Println(dog)
}
