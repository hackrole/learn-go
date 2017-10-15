package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

func main() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// the slice should now contain random bytes instead of only zeros
	fmt.Println(bytes.Equal(b, make([]byte, c)))
	fmt.Println(b)
	fmt.Printf("%+v %s", b, b)
}
