package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	// the select will not block
	// umount this to try
	c <- 2
	select {
	case i := <-c:
		fmt.Println("c", i)
	default:
		fmt.Println("defualt")
	}
}
