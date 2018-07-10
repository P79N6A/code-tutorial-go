package main

import (
        "fmt"
        "strings"
)
/*

Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
*/
func reverseWords(s string) string {
        ss := strings.Split(s," ")
        nss := make([]string,0)
        for _,sss := range ss {
                nss = append(nss,reverse(sss))
        }
        return strings.Join(nss," ")
}
func reverse(s string) string {
        ss := ""
        for i:=len(s)-1;i>=0;i-- {
                ss += string(s[i])
        }
        return ss
}

func main() {
        fmt.Println(reverseWords("Let's take LeetCode contest"))
}
