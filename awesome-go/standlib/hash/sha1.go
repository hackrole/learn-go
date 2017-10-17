package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	io.WriteString(h, "his money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")

	fmt.Printf("%x\n", h.Sum(nil))

	data := []byte("this page intentionally left blank.")
	fmt.Printf("%x\n", sha1.Sum(data))
}
