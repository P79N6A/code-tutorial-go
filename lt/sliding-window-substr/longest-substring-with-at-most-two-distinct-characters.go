package main

import "fmt"
/*
Given a string, find the length of the longest substring T that contains at most 2 distinct characters.

For example, Given s = “eceba”,

T is "ece" which its length is 3.
包含两个不同字符的最长长度
*/
func lengthOfLongestSubstringTwoDistinct(s string) int {
        letters := make(map[byte]int)
        max := 0
        left,right,slen,counter := 0,0,len(s),0
        for ;right < slen; {
                /*
                确定右边界
                */
                letters[s[right]] += 1
                if letters[s[right]] == 1 { // new char
                        // 如果是个新增的字符,counter + 1,counter代表字符的个数
                        counter += 1
                }
                right += 1
                // 如果字符超过3个,则需要处理左边界
                for ;counter == 3; { //这个循环判断进入的条件就是右边界确定完了
                        // 确定左边界
                        letters[s[left]] -= 1
                        if letters[s[left]] == 0 { // 可能变化的左边界
                                counter -= 1
                                // 左右边界都处理完了的一种可能
                                // left bound,这个肯定是2个的情况
                                if right - left > max {
                                        max = right - left
                                }
                        }
                        left += 1
                }
        }
        return max
}

func main() {
        fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
}
