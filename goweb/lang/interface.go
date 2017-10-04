package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s", h.name, h.phone)
}

type Men interface {
	SayHi()
}

func main() {
	mike := Student{Human{"mike", 25, "21231"}, "MIT", 0.00}
	paul := Employee{Human{"Sam", 52, "1231241"}, "GOalng Inc", 1000}

	var i Men
	i = mike
	i.SayHi()
	i = paul
	i.SayHi()
}
