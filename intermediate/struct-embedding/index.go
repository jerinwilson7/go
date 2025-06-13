package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	empId  string
	salary float64
}

func (p Person) introduce() {
	fmt.Printf("Hi I am %s and I am %d years old \n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi I am %s and employee ID is %s and i Earn %2f \n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		empId:  "1",
		salary: 54620,
		Person: Person{name: "John Doe", age: 56},
	}

	fmt.Println("Employee", emp.name) //Accessing the embedded struct field emp.person.name
	fmt.Println("Age", emp.age)       // Same as here
	fmt.Println("EMP ID", emp.empId)
	fmt.Println("NAME", emp.name)
	fmt.Println("SALARY", emp.salary)
	fmt.Println("PERSON", emp.Person)

	emp.Person.introduce()
	emp.introduce()

}
