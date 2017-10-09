package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(s []byte) (int, error) {
	for i, _ := range s {
		s[i] = 'A'
	}
	return len(s), nil
}

func main() {
	reader.Validate(MyReader{})
}
