package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("%v is %v\n", v, v*2)
	case string:
		fmt.Println("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Println("I donnot know about type %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
