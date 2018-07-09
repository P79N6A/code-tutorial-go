package main

import "fmt"

/*
Given a string array words, find the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. You may assume that each word will contain only lower case letters. If no such two words exist, return 0.

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
Explanation: No such pair of wordsj.
*/

func maxProduct(words []string) int {
    wf := make([]int,len(words))
    for i:=0;i<len(words);i++ {
        w := words[i]
        t:=0
        for j:=0;j<len(w);j++ {
            t |= 1 << (w[j]-'a')
        }
        wf[i]=t
    }
    product := 0
    for i:=0;i<len(words);i++ {
        for j:=i+1;j<len(words);j++ {
            if wf[i]&wf[j]==0 && len(words[i])*len(words[j])>product {
                product = len(words[i])*len(words[j])
            }
        }
    }
    return product
}

func main() {
    fmt.Println(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"}))
    fmt.Println(1<<25)
    fmt.Println(1<<('z'-'a'))
}
