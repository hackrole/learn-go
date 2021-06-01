package main

import (
	"embed"
	"fmt"
)

//--go:embed hello.txt
// var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
//go:embed hello2.txt
var f embed.FS

// export and unexport
//go:embed hello.txt
var s string

//go:embed hello.txt
var S string

//go:embed p
var ff embed.FS

func main() {
	fmt.Println(s)
	fmt.Println(S)
	fmt.Println(b)

	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("hello2.txt")
	fmt.Println(string(data))

	data, _ = f.ReadFile("p/hello.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("p/hello2.txt")
	fmt.Println(string(data))

	//function-scope embed
	//--go:embed hello.txt
	// var s string
	//--go:embed hello.txt
	// var s2 string
	// fmt.Println(s, s2)
}
