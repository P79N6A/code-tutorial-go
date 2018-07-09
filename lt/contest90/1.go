package main

import "fmt"
/*
856. Score of Parentheses
Given a balanced parentheses string S, compute the score of the string based on the following rule:

() has score 1
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.


Example 1:

Input: "()"
Output: 1
Example 2:

Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
Example 4:

Input: "(()(()))"
Output: 6
*/
func scoreOfParentheses(S string) int {
        return solve(S)
}
func solve(s string) int {
        if len(s) <= 0 {return 0}
        l,r := 0,0
        ret := 0
        last := 0
        for i:=0;i<len(s);i++ {
                if s[i]=='(' {l+=1}
                if s[i]==')' {r+=1}
                if l == r {
                        if l + r == 2 {
                                ret += 1
                        } else {
                                ret += 2*solve(s[(last+1):(last+l+r-1)])
                        }
                        last = i+1
                        l,r = 0,0
                }
        }
        return ret
}


func main() {
        fmt.Println(scoreOfParentheses("()"))
        fmt.Println(scoreOfParentheses("(())"))
        fmt.Println(scoreOfParentheses("(()(()))"))
}
