package main

import "fmt"

/*

Given a string which consists of lowercase or uppercase letters,
find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
*/
func longestPalindrome(s string) int {
    freq := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        freq[s[i]]+=1
    }
    num := 0
    odd := 0
    for _,v := range freq {
        if v%2 != 0 {
            odd = 1
        }
        num += v/2
    }
    return odd + 2*num
}
func main() {
    fmt.Println(longestPalindrome("abccccddd"))
    fmt.Println(longestPalindrome(""))
    fmt.Println(longestPalindrome("a"))

}
