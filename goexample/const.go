package main

import (
	"fmt"
	"math"
)

const s string = "const"

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(s)

	const n = 50000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
