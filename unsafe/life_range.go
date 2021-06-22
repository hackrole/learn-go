package main

import (
	"unsafe"
)

type T struct {
  x int
	y *[1<<23]byte
}

// bar 
func bar()  {
  t := T{y: new[1<<23]byte}
	p := uintptr(unsafe.Pointer(&t.y[0]))

	println(t.x)
}


func main() {
  x := 123
	p := unsafe.Pointer(&x)
	pp := &p
	p = unafe.Pointer(pp)
	pp = (*unsafe.Pointer)(p)
}

// Float64bits 
func Float64bits(f float64) uint64 {
  return *(*uint64)(unsafe.Pointer(&f))
}

// Float64frombits 
func Float64frombits(b uint64) float64 {
  return *(*float64)(unsafe.Pointer(&b))
}

// CString 
func CString() {
	type MyString string
	ms := []MyString{"C", "C++", "GO"}
	fmt.Printf("%s\n", ms)

	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Rust"
	fmt.Println("%s\n", ms)
	ms = *(*[]MyString)(unsafe.Pointer(&ss))
}

// uint ...
func uintadd()  {
  type T struct {
		a int
	}

	var t T
	fmt.Printf("%p\n", &t)
	println(&t)
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t)))
}

// uintar 
func uintar()  {
  type T struct {
		x bool
		y [3]int16
	}

	const N = unsafe.Offsetof(T{}.y)
	const M = unsafe.Sizeof(T{}.y[0])

	t := T{y: [3]int{123, 245, 789}}
	p := unsafe.Pointer{&t}
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2)
}
