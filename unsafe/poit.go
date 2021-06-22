package main

import (
	"fmt"
	"unsafe"
)

// createInt ...
func createInt() *int {
	return new(int)
}

// foo
func foo() {
	p0, y, z = createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y)
	var p2 = uintptr(unsafe.Pointer(z))

	p2 += 2
	p2--
	p2--
	*p0 = 1
	*(*int)(p1) = 2
	*(*int)(unsafe.Pointer(p2)) = 3
}

func main() {
	fmt.Println("hello world")
}
