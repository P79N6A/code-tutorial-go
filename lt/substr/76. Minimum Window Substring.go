package main

import "fmt"
/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
        letters := make(map[byte]int)
        for _,tt := range t {
                letters[byte(tt)] += 1
        }
        counter := len(letters)
        left,right,slen,tlen := 0,0,len(s),len(t)
        min,minstr := slen,""
        for ;right<slen; {
                if _,ok := letters[s[right]];ok {
                        letters[s[right]] -= 1
                        if letters[s[right]] == 0 {
                                counter -= 1
                        }
                }
                right += 1
                for ;counter ==0; {
                        if _,ok := letters[s[left]];ok {
                                letters[s[left]] += 1
                                if letters[s[left]] > 0 {
                                        counter += 1
                                }
                                if right - left >= tlen && right-left <= min {
                                        minstr = s[left:right]
                                        min = right-left
                                }
                        }
                        left += 1
                }

        }
        return minstr
}

func main() {
        fmt.Println(minWindow("a","a"))
        //fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}
