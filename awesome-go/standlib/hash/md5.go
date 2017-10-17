package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	h := md5.New()
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, "The fog is getting thicker!")
	fmt.Printf("%x\n", h.Sum(nil))
	io.WriteString(h, "And leon's getting laaarger!")
	fmt.Printf("%x\nb", h.Sum(nil))

	f, err := os.Open("./example.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h = md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))

	data := []byte("these pretzels are making me thirsty.")
	fmt.Printf("%x\n", md5.Sum(data))
}
