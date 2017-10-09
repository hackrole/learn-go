package main

import (
	"fmt"
)

func fib() func() int {
	index := 0
	first, second := 0, 1
	return func() int {
		index += 1
		if index == 1 {
			return first
		} else if index == 2 {
			return second
		} else {
			first, second = second, first+second
			return second
		}
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
