package main

import (
        "fmt"
)
/*
Given a string array words, find the maximum value of length(word[i]) * length(word[j])
 where the two words do not share common letters. You may assume that each word will contain only lower case letters. If no such two words exist, return 0.

Example 1:

Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16
Explanation: The two words can be "abcw", "xtfn".
Example 2:

Input: ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4
Explanation: The two words can be "ab", "cd".
Example 3:

Input: ["a","aa","aaa","aaaa"]
Output: 0
Explanation: No such pair of words.
*/
func maxProduct(words []string) int {
        product := 0
        for i:=0;i<len(words);i++ {
                for j:=i+1;j<len(words);j++ {
                        if !hascommon(words[i],words[j]) {
                                p := len(words[i]) * len(words[j])
                                if p > product {product = p}
                        }
                }
        }
        return product
}
func hascommon(w1,w2 string) bool {
        ws := make(map[byte]bool)
        for i:=0;i<len(w1);i++ {
                ws[w1[i]]=true
        }
        for i:=0;i<len(w2);i++ {
                if ws[w2[i]]==true {
                        return true
                }
        }
        return false
}

func main() {
        fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}))
}
