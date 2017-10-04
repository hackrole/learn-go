package main

import (
	"fmt"
	"math"
	"reflect"
)

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rect{12, 2}
	r2 := Rect{9, 5}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println(r1.area(), r2.area(), c1.area(), c2.area())

	fmt.Println(reflect.TypeOf(c2))
	fmt.Printf("%v %#v %T %%", c2, c2, c2, c2)

	a := "hello world"
	fmt.Printf("%x %X", a, a)
}
