package main

import "fmt"

/*
Given a string S and a string T,
find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
    slen,tlen := len(s),len(t)
    letters := make(map[byte]int)
    for i:=0;i<tlen;i++ {
        letters[t[i]] += 1
    }
    fmt.Println(letters)
    substr,sublen := "",slen
    l,r,counter := 0,0,len(letters)
    for ;r<slen; {
        if _,ok := letters[s[r]];ok {
            letters[s[r]] -= 1
            if letters[s[r]] == 0 {
                counter -= 1
            }
        }
        r += 1
        for ;counter==0; {
            if _,ok := letters[s[l]];ok {
                letters[s[l]] += 1
                if letters[s[l]] > 0 {
                    counter += 1
                }
            }
            if r - l <= sublen {
                sublen = r - l
                substr = s[l:r]
            }
            fmt.Println(l,r,sublen,substr)
            l += 1
        }
    }
    return substr
}

func main() {
    //fmt.Println(minWindow("ADOBECODEBANC","ABC"))
    fmt.Println(minWindow("cabwefgewcwaefgcf","cae"))
}
