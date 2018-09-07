package main

import "fmt"

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "*"
Output: true
Explanation: '*' matches any sequence.
Example 3:

Input:
s = "cb"
p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
Example 4:

Input:
s = "adceb"
p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
Example 5:

Input:
s = "acdcb"
p = "a*c?b"
Output: false

*/

func main() {
    //fmt.Println(isMatch("adceb","*a*b"))
    //fmt.Println(isMatch("acdcb","a*c?b"))
    //fmt.Println(isMatch("aa","*"))
    //fmt.Println(isMatch("ho","ho**"))
    fmt.Println(isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab","b*b*ab**ba*b**b***bba"))

}

func isMatch(s string, p string) bool {
    return solve(s,p,make(map[string]bool))
}
func solve(s string, p string,cache map[string]bool) bool {
    key := p + "_" + s
    if _,ok := cache[key];ok {
        return cache[key]
    }
    if p == "" {
        return s == ""
    }
    if s == "" {
        if p[0]=='*' {
            return solve(s,p[1:],cache)
        }
        cache[key]=false
        return false
    }
    if p[0] == '?' {
        return solve(s[1:],p[1:],cache)
    } else if p[0] == '*' {
        for i:=0;i<=len(s);i++ {
            if solve(s[i:],p[1:],cache) {
                cache[key]=true
                return true
            }
        }
        cache[key]=false
        return false
    } else {
        if s[0]==p[0] {
            return solve(s[1:],p[1:],cache)
        }
        cache[key]=false
        return false
    }
    cache[key]=false
    return false
}