package main

import (
    "math"
    "strings"
    "fmt"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func longestCommonPrefix(strs []string) string {
    mstr,mlen := "",math.MaxInt8
    for _,s := range strs {
        if len(s) < mlen {
            mstr,mlen = s,len(s)
        }
    }
    if mstr == "" {return ""}
    l,r := 0,mlen
    for l < r {
        m := (l+r)/2
        md := mstr[:m+1]
        prefix := true
        for _,s := range strs {
            if !strings.HasPrefix(s,md) {
                prefix = false
                break
            }
        }
        fmt.Println(md,l,r,m,prefix)
        if prefix {
            l = m + 1
        } else {
            r = m
        }
    }
    return mstr[:r]
}

func main() {
//    fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
    fmt.Println(longestCommonPrefix([]string{"abab","aba","abc"}))

}
