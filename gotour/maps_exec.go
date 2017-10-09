package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	fmt.Println(words)
	results := make(map[string]int)
	for _, w := range words {
		if _, ok := results[w]; ok {
			results[w] += 1
		} else {
			results[w] = 1
		}
	}

	fmt.Println(results)
	return results
}

func main() {
	wc.Test(WordCount)
}
