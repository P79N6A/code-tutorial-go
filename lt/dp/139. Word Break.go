package main

import (
    "strings"
    "fmt"
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
func wordBreak2(s string, wordDict []string) bool {
    /*
    动态规划
    dp[i]=dp[j]&&s[i-j] in worddict
    典型区间模型，看下实现。。
    */
    //
    maxl := 0
    wd := make(map[string]bool)
    for _,w := range wordDict {
        if len(w) > maxl {
            maxl = len(w)
        }
        wd[w]=true
    }
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := i - 1; i-j<=maxl&&j >= 0; j-- {
            if _,ok := wd[s[j:i]];ok&&dp[j]{
                dp[i] = dp[j]
            }
        }
    }
    return dp[len(s)]
}
func wordBreakR(s string, wordDict []string) bool {
    if len(s) == 0 {
        return true
    }
    for _, w := range wordDict {
        if strings.HasPrefix(s, w) {
            b := wordBreakR(s[len(w):], wordDict)
            if b {
                return true
            }
        }
    }
    return false
}
func wordBreak(s string, wordDict []string) bool {
    wSet := make(map[byte]bool)
    for _, w := range wordDict {
        for i := 0; i < len(w); i++ {
            wSet[w[i]] = true
        }
    }
    for i := 0; i < len(s); i++ {
        if _,ok := wSet[s[i]];!ok {
            return false
        }
    }
    return wordBreakR(s,wordDict)
}

func main() {
    fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))

}
