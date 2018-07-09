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
    for i := 0; i < len(t); i++ {
        letters[t[i]] += 1 // hash记录目标词频
    }
    counter := len(letters)
    left, right, slen := 0, 0, len(s)
    min, minstr := slen, ""
    for right < slen {
        if _, ok := letters[s[right]]; ok {
            letters[s[right]] -= 1
            if letters[s[right]] == 0 {
                // 如果有一个字符都没有，counter-1
                // 这个是== 而不是<= 因为letters[i]可能是负数
                counter -= 1
            }
        }
        right += 1
        for counter == 0 { // 说明所有的字符都出现过了，确定了右边界right
            if _, ok := letters[s[left]]; ok {
                letters[s[left]] += 1
                if letters[s[left]] > 0 { // 某一个字符多出来了，说明slidingwindow缺少这个字符，左边界确定
                    counter += 1
                }
                if right-left <= min {
                    // 一个可能的结果；只要确认过右边界，再确认左边界之前，都可能是个有效结果
                    minstr = s[left:right]
                    min = right - left
                }
            }
            left += 1
        }
    }
    return minstr
}

func main() {
    fmt.Println(minWindow("a", "a"))
    //fmt.Println(minWindow("ADOBECODEBANC","ABC"))
}
