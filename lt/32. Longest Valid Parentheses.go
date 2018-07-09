package main

import "fmt"
/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/
func longestValidParentheses(s string) int {
        // 所有的dp,都需要有最前边的空位值,否则会有越界问题
        s = "x" + s
        slen := len(s)
        dp := make([]int,slen)
        result := 0
        for i:=1;i<slen;i++ {
                // dp[i] 表示以i结尾最长的串长度
                // 算法,如果当前是) 则看和上一个为结尾形成的dp[i-1]的前边字符是否形成串
                if s[i] == ')' {
                        if s[i-1-dp[i-1]] == '(' {
                                dp[i] += dp[i-1] + 2
                        }
                        dp[i] += dp[i-dp[i]]
                        if result < dp[i] {
                                result = dp[i]
                        }
                }
        }
        return result
}

func main() {
        //fmt.Println(longestValidParentheses(")()())"))
        fmt.Println(longestValidParentheses("()"))
}
