package main

import "fmt"
/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

Input: "hello"
Output: "holle"
Example 2:

Input: "leetcode"
Output: "leotcede"
Note:
The vowels does not include the letter "y".
*/

func main() {
        fmt.Println(reverseVowels("hello"))
}

func reverseVowels(ss string) string {
        l,r := 0,len(ss)-1
        s := []byte(ss)
        for l<r {
                if !isVowel(s[l]) {
                        l += 1
                        continue
                }
                if !isVowel(s[r]) {
                        r -= 1
                        continue
                }
                s[l],s[r]=s[r],s[l]
                l += 1
                r -= 1
        }
        return string(s)
}
var vowel = map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true,'A':true,'E':true,'I':true,'O':true,'U':true}
func isVowel(b byte) bool {
        return vowel[b]
        //return map[byte]bool{'a':true,'e':true,'i':true,'o':true,'u':true,'A':true,'E':true,'I':true,'O':true,'U':true}[b] == true
}