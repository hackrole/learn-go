package main

import (
	"fmt"

	"github.com/derekparker/trie"
)

func main() {
	t := trie.New()

	// add can take in meta info which can be stored with key
	// the paricular key
	t.Add("foobar", 1)

	// find a key with
	node, _ := t.Find("foobar")
	meta := node.Meta()
	// use meta use meta.(type)
	fmt.Printf("%+V %T", meta, meta)

	// remove keys
	t.Remove("foobar")

	// prefix search
	t.PrefixSearch("foo")

	// fast test valid prefix
	t.HasKeysWithPrefix("foo")

	// fuzzy search
	t.FuzzySearch("fb")
}
