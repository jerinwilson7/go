package main

import "fmt"

type Shape struct {
	Rectangle //struct embedding
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(i float64) {
	r.width *= i
	r.length *= i
}

type MyInt int

// Method on a user defined type

// in-order to associate a method with type or struct we need to mention the type  before the method
func (m MyInt) isPositive() bool {
	if m > 0 {
		return true
	} else {
		return false
	}
}

func (MyInt) WelcomeMessage() string {
	return "Welcome to MyInt"
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()

	fmt.Println(area)

	rect.Scale(2)
	rect.Scale(5)
	rect.Area()
	rect.Scale(4)
	rect.length = 20
	rect.Area()

	fmt.Println(rect)

	num1 := MyInt(2)
	num2 := MyInt(-3)

	fmt.Println(num1.isPositive())
	fmt.Println(num2.isPositive())
	fmt.Println(num1.WelcomeMessage())
	fmt.Println(MyInt.WelcomeMessage(1))

	s := Shape{Rectangle: Rectangle{length: 40, width: 70}}
	s.Area()
	s.Rectangle.Area()
	fmt.Println(s)
}
