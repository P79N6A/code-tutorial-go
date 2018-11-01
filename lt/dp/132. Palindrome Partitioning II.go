package main

import "fmt"

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example:

Input: "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/

func main() {
    fmt.Println(minCut("aaadcccb"))
    fmt.Println(minCut("aaa"))
}

func minCut(s string) int {
    return minCutF(s,make(map[string]int))
}
func minCutF(s string,cache map[string]int) int {
    if len(s)<=1 {return 0}
    if _,ok := cache[s];ok {
        return cache[s]
    }
    if isPalindrome(s) {
        cache[s]=0
        return 0
    }
    ans := len(s)+1
    for i:=1;i<len(s);i++ {
        if isPalindrome(s[:i]) {
            ans = min(ans,1+minCutF(s[i:],cache))
        }
    }
    cache[s]=ans
    return ans
}
func min(a,b int) int {
    if a<b {return a}
    return b
}
func isPalindrome(a string) bool {
    l,r := 0,len(a)-1
    for l<r {
        if a[l]!=a[r] {return false}
        l+=1
        r-=1
    }
    return true
}
