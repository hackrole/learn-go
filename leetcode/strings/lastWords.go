package main

import (
	"fmt"
	"strings"
)

func lenOfLastWord(word string) int {
	words := strings.Fields(word)
	last_word := words[len(words)-1]
	return len(last_word)
}

func main() {
	fmt.Println(lenOfLastWord("hello world, hello golang."))
}
