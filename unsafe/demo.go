package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T struct {
		c string
	}
	type S struct {
		b bool
	}
	var x struct {
		a int64
		*S
		T
	}

	fmt.Println(unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Offsetof(x.S))
	fmt.Println(unsafe.Offsetof(x.T))

	fmt.Println(unsafe.Offsetof(x.c))
	fmt.Println(unsafe.Offsetof(x.S.b))
}
