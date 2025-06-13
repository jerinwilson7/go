package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area is : ", g.area())
	fmt.Println("Perimeter is  :", g.perim())
}

func detectCircle(g Geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("Circle with radius", c.radius)
	}
}

func detectRectangle(g Geometry) {
	if r, ok := g.(Rectangle); ok {
		fmt.Println("Rectangle with height and width of", r.height, r.width)
	}
}

func main() {

	circle1 := Circle{radius: 5}
	rectangle1 := Rectangle{height: 21, width: 25}

	measure(rectangle1)
	measure(circle1)

	detectCircle(circle1)
	detectCircle(rectangle1)

	detectRectangle(rectangle1)
	detectRectangle(circle1)

}
