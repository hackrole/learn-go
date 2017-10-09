package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(s []byte) (int, error) {
	fmt.Println(s, rot.r)
	//for i, _ := range s
	//  s[i] += 13
	//  if s[i] > 'z' {
	//    s[i] -= 26
	//  }
	//}
	return len(s), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrqgui pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
