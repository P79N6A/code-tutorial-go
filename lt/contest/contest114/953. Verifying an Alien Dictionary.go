package main

import "fmt"

/*
In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.



Example 1:

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
Example 2:

Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
Example 3:

Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).


Note:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are english lowercase letters.
*/
func main() {
    fmt.Println(isAlienSorted([]string{"hello","leetcode"},"hlabcdefgijkmnopqrstuvwxyz"))
    fmt.Println(isAlienSorted([]string{"word","world","row"},"worldabcefghijkmnpqstuvxyz"))
    fmt.Println(isAlienSorted([]string{"apple","app"},"abcdefghijklmnopqrstuvwxyz"))
    fmt.Println(isAlienSorted([]string{"kuvp","q"},"ngxlkthsjuoqcpavbfdermiywz"))
}
func isAlienSorted(words []string, order string) bool {
    orderMap := make(map[byte]byte)
    for i:=0;i<len(order);i++ {
        orderMap[order[i]]=byte(i)
    }
    transfunc := func(w string) string {
        nw := []byte{}
        for i:=0;i<len(w);i++ {
            nw = append(nw,'a'+orderMap[w[i]])
        }
        return string(nw)
    }
    nword := make([]string,0,len(words))
    for _,word := range words {
        nword = append(nword,transfunc(word))
    }
    for i:=1;i<len(nword);i++ {
        if nword[i-1]>nword[i] {
            return false
        }
    }
    return true
}