package main

import "fmt"

/*
Given a string s, you are allowed to convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.

Example 1:

Input: "aacecaaa"
Output: "aaacecaaa"
Example 2:

Input: "abcd"
Output: "dcbabcd"

*/

func main() {
    fmt.Println(shortestPalindrome("aacecaaa"))
    fmt.Println(shortestPalindrome("aa"))

}
func reverse(s string) string {
    nr := make([]byte,len(s))
    l,r := 0,len(s)-1
    for l<len(s){
        nr[l]=s[r]
        l+=1
        r-=1
    }
    return string(nr)
}
func shortestPalindrome(s string) string {
    idx:=len(s)
    for ;idx>=0;idx-- {
        if isPalindrom(s[:idx]) {
            return reverse(s[idx:])+s
        }
    }
    return ""
}
func isPalindrom(s string) bool {
    l,r := 0,len(s)-1
    for l<r {
        if s[l]!=s[r]{return false}
        l+=1
        r-=1
    }
    return true
}