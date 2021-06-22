// need go version >= 1.17
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)

	eleSize := int(unsafe.Sizeof(a[0]))
	p9 := &a[9]
	up9 := unsafe.Pointer(p9)
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println(*p3)
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	_ = unsafe.Add(up9, 7*eleSize)
	_ = unsafe.Slice(p9, 8)
}
