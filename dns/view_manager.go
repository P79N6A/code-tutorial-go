package main

import (
	"fmt"
	"mustard/base/container"
)
type E struct {
	name string
	age int
}

func main() {
	trie := container.NewTrie()
	trie.Insert([]byte("01234"))
	fmt.Println(trie.Root.DumpChild())
	fmt.Println(trie.IsPrefix([]byte("01234")))
	fmt.Println(trie.IsPrefix("012345"))
	fmt.Println(trie.IsPrefix("0123"))
}
