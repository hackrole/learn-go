package main

import (
	"fmt"

	"github.com/zoumo/goset"
)

func main() {
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.0, 2.0, 3.0}
	strings := []string{"1", "2", "3"}

	fmt.Printf("%+V %T\n", goset.NewSetFrom(ints))
	fmt.Printf("%+V %T\n", goset.NewSafeSetFrom(ints))
	fmt.Printf("%+V %T\n", goset.NewSetFrom(floats))
	fmt.Printf("%+V %T\n", goset.NewSafeSetFrom(floats))
	fmt.Printf("%+V %T\n", goset.NewSetFrom(strings))
	fmt.Printf("%+V %T\n", goset.NewSafeSetFrom(strings))
}
