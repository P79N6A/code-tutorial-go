package main

import "fmt"
/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring2(s string) int {
        left, right, slen := 0, 0, len(s)
        letters := make(map[byte]int)
        counter := 0
        max, maxsub := 0, ""

        for ; right < slen; {
                letters[s[right]] += 1
                if letters[s[right]] > 1 {
                        counter += 1
                }
                right += 1
                for ; counter == 1; {
                        letters[s[left]] -= 1
                        if letters[s[left]] == 1 {
                                counter -= 1
                                //
                        }
                        left += 1
                }
                if counter == 0 {
                        if right - left >= max {
                                max = right - left
                                maxsub = s[left:right]
                        }
                }
        }
        fmt.Println(maxsub, left, right)
        return max
}
func lengthOfLongestSubstring(s string) int {
        left, right, slen := 0, 0, len(s)
        letters := make(map[byte]int)
        max, maxsub := 0, ""
        for ; right < slen; {
                if letters[s[right]] <= 0 {
                        letters[s[right]] += 1
                        right += 1
                } else {
                        //find right bound
                        if right - left >= max {
                                max = right - left
                                maxsub = s[left:right]
                        }
                        letters[s[left]] -= 1
                        left += 1
                }
        }
        if right - left >= max {
                max = right - left
                maxsub = s[left:right]
        }
        fmt.Println(maxsub, left, right)
        return max
}

func main() {
        ///        fmt.Println(lengthOfLongestSubstring("p"))
        //fmt.Println(lengthOfLongestSubstring("p"))
        //fmt.Println(lengthOfLongestSubstring2("p"))

        fmt.Println(lengthOfLongestSubstring("pwwkew"))
        fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}
