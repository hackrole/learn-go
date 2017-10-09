package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// cannot set methods on build-in type
//func (f float64) Abs() float64 {
//  return float64(f)
//}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	//var a float64 = 3.12
	//fmt.Println(a.Abs())
}
