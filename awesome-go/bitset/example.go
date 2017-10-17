package main

import (
	"fmt"
	"math/rand"

	"github.com/willf/bitset"
)

func main() {
	fmt.Printf("Hello from BItSet!\n")

	var b bitset.BitSet
	// play some go fish
	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52))
		card2 := uint(rand.Intn(52))

		b.Set(card1)
		if b.Test(card2) {
			fmt.Println("Go Fish!")
		}
		b.Clear(card1)
	}

	// chaining
	b.Set(10).Set(11)

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("the following bit is set:", i)
	}
	if b.Intersection(bitset.New(100).Set(100)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesnot work.")
	}
}
