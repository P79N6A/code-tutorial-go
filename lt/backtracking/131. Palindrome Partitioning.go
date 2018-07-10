package main

import "fmt"

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]

*/
func partition(s string) [][]string {
    if len(s) <= 0 {return nil}
    if len(s) == 1 {return [][]string{[]string{s}}}
    ret := make([][]string,0)
    for i:=len(s)-1;i>=0;i-- {
        if valid(s[i:]) {
            sub := partition(s[:i])
            for _,ss := range sub {
                ss = append(ss,s[i:])
                ret = append(ret,ss)
            }
        }
    }
    if valid(s) {
        ret = append(ret,[]string{s})
    }
    return ret
}
func valid(s string) bool {
    l,r := 0,len(s)-1
    for l < r && s[l] == s[r] {
        l,r = l+1,r-1
    }
    return r-l <= 0
}

func main() {
    fmt.Println(partition("aab"))
    fmt.Println(partition("aa"))

}
