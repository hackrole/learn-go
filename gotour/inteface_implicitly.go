package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (t T) P() {
	fmt.Println("t.s:", t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
	// this would error
	//i.P()
}
