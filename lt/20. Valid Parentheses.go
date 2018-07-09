package main

import "fmt"
/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

*/

func isValid(s string) bool {
        st := make([]byte,len(s))
        p := 0
        for i:=0;i<len(s);i++ {
                if s[i] == '(' || s[i] == '{' || s[i] == '[' {
                        st[p]=s[i]
                        p += 1
                } else {
                        if p <= 0 {
                                return false
                        }
                        if s[i] == ')' && st[p-1] == '(' {
                                p -= 1
                        } else if s[i] == ']' && st[p-1] == '[' {
                                p -= 1
                        } else if s[i] == '}' && st[p-1] == '{' {
                                p -= 1
                        } else {
                                return false
                        }
                }
        }
        return p == 0
}

func main() {
        fmt.Println(isValid("{]}"))
}
