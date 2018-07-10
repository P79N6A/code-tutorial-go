package main

import "fmt"
/*
Given a list of words, we may encode it by writing a reference string S and a list of indexes A.

For example, if the list of words is ["time", "me", "bell"], we can write it as S = "time#bell#" and indexes = [0, 2, 5].

Then for each index, we will recover the word by reading from the reference string from that index until we reach a "#" character.

What is the length of the shortest reference string S possible that encodes the given words?

Example:

Input: words = ["time", "me", "bell"]
Output: 10
Explanation: S = "time#bell#" and indexes = [0, 2, 5].
Note:

1 <= words.length <= 2000.
1 <= words[i].length <= 7.
Each word has only lowercase letters.
*/

func minimumLengthEncoding(words []string) int {
        parent := make(map[string]string)
        for _,word := range words {
                parent[word] = ""
        }
        for _,word := range words {
                for i:=1;i<len(word);i++ {
                        if _,ok := parent[word[i:]];ok {
                                parent[word[i:]] = word
                        }
                }
        }
        fmt.Println(parent)
        rparent := make(map[string]bool)
        for k,v := range parent {
                if v == "" {
                        rparent[k]=true
                } else {
                        rparent[v]=true
                }
        }
        fmt.Println(rparent)
        l := 0
        for k,_ := range rparent {
                l += len(k) + 1
        }
        return l
}
func main() {
        fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
        
}
