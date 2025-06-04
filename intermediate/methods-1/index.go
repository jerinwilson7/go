package main

import "fmt"

type Rect struct {
	height int
	width  int
}

func (r *Rect) area() int {
	return r.height * r.width
}

func (r Rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {

	r := Rect{height: 20, width: 24}

	fmt.Println(r)

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rc := &r

	fmt.Println(rc)

	rc.height = 10

	fmt.Println(r)
	fmt.Println(rc)

}
