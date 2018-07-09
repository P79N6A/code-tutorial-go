package main

import (
    "strconv"
    "strings"
    "fmt"
)

/*

Given an encoded string, return it's decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Examples:

s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
*/
func decodeString(s string) string {
    stack := make([]byte,0)
    for i:=0;i<len(s);i++ {
        b := s[i]
        if b == ']' {
            // pop and caculate
            str,num := "",""
            for len(stack) > 0{
                tt := stack[len(stack)-1]
                if tt == '[' {
                    // get the number
                    stack = stack[:len(stack)-1]
                    for len(stack) > 0 {
                        t := stack[len(stack)-1]
                        if t >= '0' && t <= '9' {
                            num = string(t)+num
                            stack = stack[:len(stack)-1]
                        } else {
                            break
                        }
                    }
                    break
                } else {
                    str = string(tt)+str
                    stack = stack[:len(stack)-1]
                }
            }
            n,_ := strconv.Atoi(num)
            nstr := strings.Repeat(str,n)
            stack = append(stack,[]byte(nstr)...)
        } else {
            stack = append(stack,b)
        }
    }
    return string(stack)
}
func main() {
    fmt.Println(decodeString("3[a]2[bc]"))
    fmt.Println(decodeString("2[abc]3[cd]ef"))
}
