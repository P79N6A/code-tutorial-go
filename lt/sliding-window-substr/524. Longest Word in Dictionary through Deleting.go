package main

import (
    "fmt"
    "sort"
)

/*
Given a string and a string dictionary, find the longest string in the dictionary
that can be formed by deleting some characters of the given string.
If there are more than one possible results, return the longest word
with the smallest lexicographical order. If there is no possible result, return the empty string.

Example 1:
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
Example 2:
Input:
s = "abpcplea", d = ["a","b","c"]

Output:
"a"
Note:
All the strings in the input will only contain lower-case letters.
The size of the dictionary won't exceed 1,000.
The length of all the strings in the input won't exceed 1,000.
*/
func isSubsequence(s string, t string) bool {
    slen := len(s)
    if slen <= 0 {return true}
    for i,idx := 0,0;i<len(t);i++ {
        if s[idx] == t[i] {
            idx += 1
        }
        if idx >= slen {
            return true
        }
    }
    return false
}
func findLongestWord(s string, d []string) string {
    lw := ""
    sort.Strings(d)
    for _,w := range d {
        if isSubsequence(w,s) {
            if len(lw) == 0 || len(lw) < len(w) {
                lw = w
            }
        }
    }
    return lw
}
func main() {
    fmt.Println(findLongestWord("abpcplea",[]string{"ale","apple","monkey","plea"}))
    
}
