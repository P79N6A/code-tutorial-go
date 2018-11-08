package main

import "fmt"

/*
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
*/

func main() {
    t := Constructor()
    t.Insert("apple")
    fmt.Println(t.Search("apple"))
    fmt.Println(t.Search("app"))
    fmt.Println(t.StartsWith("app"))
    t.Insert("app")
    fmt.Println(t.Search("app"))

}
type TrieNode struct {
    isword bool
    next []*TrieNode // a-z
}
func NewTrieNode() *TrieNode {
    return &TrieNode{
        isword:false,
        next:make([]*TrieNode,26), // a-z
    }
}
func insert(root *TrieNode, str string) {
    if str == "" {
        root.isword=true
        return
    }
    if root.next[str[0]-'a']==nil {
        root.next[str[0]-'a']=NewTrieNode()
    }
    insert(root.next[str[0]-'a'],str[1:])
}
func startwith(root *TrieNode,str string) bool {
    if str == "" {return true}
    if root.next[str[0]-'a'] == nil {return false}
    return startwith(root.next[str[0]-'a'],str[1:])
}
func search(root *TrieNode,str string) bool {
    if str == "" { return root.isword}
    if root.next[str[0]-'a'] == nil {return false}
    return search(root.next[str[0]-'a'],str[1:])
}

type Trie struct {
    root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        root:NewTrieNode(),
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    insert(this.root,word)
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    return search(this.root,word)
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    return startwith(this.root,prefix)
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */