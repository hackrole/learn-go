package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	var p person
	p.name = "Astaxie"
	p.age = 25
	fmt.Printf("%d %s", p.age, p.name)

	p2 := person{"Tom", 25}
	p3 := person{age: 24, name: "Tom"}

	fmt.Println(p2, p3)

	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println(mark)
}
