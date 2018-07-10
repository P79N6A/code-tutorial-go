package main

import "fmt"
/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

*/
func findAnagrams(s string, p string) []int {
        // how many times show in p
        letters := make(map[byte]int)
        for _,pp := range p {
                letters[byte(pp)] += 1
        }
        counter := len(letters)
        begin,end,slen,plen := 0,0,len(s),len(p)
        result := make([]int,0)
        for ;end < slen; {
                if _,ok := letters[s[end]];ok {
                        letters[s[end]] -= 1
                        if letters[s[end]] == 0 {
                                counter -= 1
                        }
                }
                end += 1
                // find right bound
                for ;counter == 0;{
                        if _,ok := letters[s[begin]];ok {
                                letters[s[begin]] += 1
                                if letters[s[begin]] > 0 {
                                        // find left bound
                                        counter += 1
                                }
                                if end-begin == plen {
                                        result = append(result,begin)
                                }
                        }
                        begin += 1
                }

        }
        return result

}

func main() {
        fmt.Println(findAnagrams("cbaebabacd","abc"))
}
