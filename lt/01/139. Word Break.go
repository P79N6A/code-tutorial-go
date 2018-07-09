package main

import (
        "strings"
        "fmt"
        "sort"
)
/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false

*/
/*
多重背包
dp[i][w] = dp[i-1][w] || dp[i-1][w-word]

*/
// 递归
func wordBreak1(s string, wordDict []string) bool {
        values := make([]string,0)
        num := 0
        splitR(s,wordDict,values,&num)
        fmt.Println(values,num)
        return num>0
}
func splitR(s string,wordDict []string,values []string,num *int) {
        if len(s) == 0 {
                fmt.Println(values)
                *num += 1
                return
        }

        for _,word := range wordDict {
                if strings.HasPrefix(s,word) {
                        values = append(values,word)
                        splitR(s[len(word):],wordDict,values,num)
                        values = values[:len(values)-1]
                }
        }
        return
}
type SString []string
func (s SString)Swap(i,j int) {
        s[i],s[j]=s[j],s[i]
}
func (s SString)Less(i,j int) bool {
        return len(s[i]) > len(s[j])
}
func (s SString)Len() int {
        return len(s)
}

func wordBreak(s string, wordDict []string) bool {
        wset := make(map[byte]bool)
        for _,w := range wordDict {
                for i:=0;i<len(w);i++ {
                        wset[w[i]]=true
                }
        }
        for i:=0;i<len(s);i++ {
                if wset[s[i]] == false {
                        return false
                }
        }
        sort.Sort(SString(wordDict))
        fmt.Println(wordDict)
        //超时问题严重,试着加一个缓存~~~
        cache := make(map[int]bool)

        ok := split(s,wordDict,cache)
        fmt.Println(cache)
        return ok
}

func split(s string,wordDict []string,cache map[int]bool) bool {
        if len(s) == 0 {
                return true
        }
        for _,word := range wordDict {
                if strings.HasPrefix(s,word) {
                        if cache[len(s)-len(word)] == true {
                                continue
                        }
                        ok := split(s[len(word):],wordDict,cache)
                        if ok {
                                return true
                        } else {
                                cache[len(s)-len(word)]=true
                        }
                }
        }
        return false
}
func main() {
        fmt.Println(wordBreak("applepenapple",[]string{"apple","pen"}))
        //fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa","ba"}))
        fmt.Println(wordBreak("aabbcc",[]string{"aa","aaaa","bb","c","cc"}))
        //fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))
        fmt.Println(wordBreak("catsandog",[]string{"cats", "dog", "sand", "and", "cat"}))
}
