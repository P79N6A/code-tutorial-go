package main

import (
        "fmt"
        "strings"
)
/*

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]
Example 2:

Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.
Example 3:

Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
*/
func wordBreak2(s string, wordDict []string) []string {
        wd := make(map[string]bool)
        for _,w := range wordDict {
                wd[w]=true
        }
        res := make([]string,0)
        ret := make([]string,0)
        solve2(s,wd,&res,&ret)
        for _,r:= range ret {fmt.Println(r)}
        return ret
}
func solve2(s string, wd map[string]bool,res, ret *[]string) {
        if len(s) <= 0 {
                *ret = append(*ret,strings.Join(*res," "))
                return
        }
        for i:=0;i<=len(s);i++ {
                if wd[s[:i]] == true {
                        *res = append(*res, s[:i])
                        solve2(s[i:], wd, res, ret)
                        *res = (*res)[:len(*res) - 1]
                }
        }
}
func wordBreak(s string, wordDict []string) []string {
        wd := make(map[string]bool)
        for _, w := range wordDict {
                wd[w] = true
        }
        // 首先要判断是否有解,这个dp可以搞定!!
        dp := make([]bool,len(s)+1)
        dp[0]=true
        for i:=1;i<=len(s);i++ {
                for j:=i;j>=0;j-- {
                        if wd[s[j:i]]==true && dp[j] {
                                dp[i]=true
                                break
                        }
                }
        }
        if dp[len(s)]==false {
                return nil
        }

        cache := make(map[string][]string)
        return solve(s,wd,cache)
}
func solve(s string,wd map[string]bool,cache map[string][]string) []string {
        if len(cache[s]) > 0 {
                return cache[s]
        }
        if len(s)<=0{
                return nil
        }
        res := make([]string,0)
        for i := 0; i <=len(s); i++ {
                if wd[s[:i]] == true {
                        l := solve(s[i:],wd,cache)
                        if l == nil {
                                res = append(res,s[:i])
                        } else {
                                for _,ll := range l {
                                        res = append(res,s[:i]+ " " + ll)
                                }
                        }
                }
        }
        cache[s]=res
        return res
}
func main() {
        //fmt.Println(wordBreak("catsanddog",[]string{"cat", "cats", "and", "sand", "dog"}))
        fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))
}
