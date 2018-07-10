package main

import (
        "strings"
        "fmt"
)
/*
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word
in words exactly once and without any intervening characters.

Example 1:

Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
Example 2:

Input:
  s = "wordgoodstudentgoodword",
  words = ["word","student"]
Output: []
*/
func findSubstring(s string, words []string) []int {
        wmap := make(map[string]int)
        for _,w := range words {
                wmap[w] += 1
        }
        rmap := make(map[string]int)
        for k,v := range wmap {rmap[k]=v}
        res := make([]int,0)
        counter := len(words)
        l,r,slen := 0,0,len(s)
        for ;r<slen; {
                match := false
                for w,_ := range wmap {
                        if strings.HasPrefix(s[r:], w) {
                                wmap[w] -= 1

                                if wmap[w] >= 0 {
                                        match = true
                                        counter -= 1
                                }
                                r += len(w)
                                break
                        }
                }
                for ;counter==0&&match ; {
                        for k,v := range rmap {wmap[k]=v}
                        counter = len(words)
                        res = append(res,l)
                        l += 1
                        r = l
                }
                if match == false {
                        for k,v := range rmap {wmap[k]=v}
                        counter = len(words)
                        l += 1
                        r = l
                }
        }
        return  res
}

func main() {
        fmt.Println(findSubstring("wordgoodgoodgoodbestword",[]string{"word","good","best","word"}))
        //fmt.Println(findSubstring("barfoofoobarthefoobarman",[]string{"bar","foo","the"}))
}
