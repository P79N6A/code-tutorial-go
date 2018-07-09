package main

import "fmt"

/*

Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
In other words, one of the first string's permutations is the substring of the second string.
Example 1:
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False
Note:
The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].

*/
func checkInclusion(s1 string, s2 string) bool {
    freq := make(map[byte]int)
    for i:=0;i<len(s1);i++ {
        freq[s1[i]]+=1
    }
    miss := make(map[byte]int)
    l,r := 0,0
    counter := len(freq)
    for r < len(s2) {
        if freq[s2[r]] > 0 {
            freq[s2[r]] -= 1
            if freq[s2[r]] == 0 {
                counter -= 1
            }
        } else {
            miss[s2[r]]+=1
        }
        r += 1
        fmt.Println(freq,miss,counter,l,r)
        for counter == 0 && l <= r {
            if len(miss) == 0 && counter == 0 {
                return true
            }
            if miss[s2[l]] > 0 {
                miss[s2[l]] -= 1
                if miss[s2[l]] == 0 {
                    delete(miss,s2[l])
                }
            } else if _,ok := freq[s2[l]];ok {
                freq[s2[l]]+=1
                if freq[s2[l]]==1 {
                    counter += 1
                }
            }
            //fmt.Println(miss,counter,freq)
            l += 1
        }
    }
    return len(miss) == 0 && counter==0
}

func main() {
//    fmt.Println(checkInclusion("a","ab"))
    fmt.Println(checkInclusion("ab","a"))
//    fmt.Println(checkInclusion("adc","dcda"))
//    fmt.Println(checkInclusion("ab","eidboaoo"))

}
