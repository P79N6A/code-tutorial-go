package main

import "fmt"

/*

S and T are strings composed of lowercase letters. In S, no letter occurs more than once.

S was sorted in some custom order previously. We want to permute the characters of T so that they match the order that S was sorted. More specifically, if x occurs before y in S, then x should occur before y in the returned string.

Return any permutation of T (as a string) that satisfies this property.

Example :
Input:
S = "cba"
T = "abcd"
Output: "cbad"
Explanation:
"a", "b", "c" appear in S, so the order of "a", "b", "c" should be "c", "b", and "a".
Since "d" does not appear in S, it can be at any position in T. "dcba", "cdba", "cbda" are also valid outputs.


Note:

S has length at most 26, and no character is repeated in S.
T has length at most 200.
S and T consist of lowercase letters only.
*/

func main() {
    fmt.Println(customSortString("cba","cbad"))
}
func customSortString(S string, T string) string {
    /*
    看T中出现次数，然后按照S中顺序输出；没出现在S的放后边给
    */
    freq := make(map[byte]int)
    ret := make([]byte,0)
    for i:=0;i<len(T);i++ {
        freq[T[i]]+=1
    }
    for i:=0;i<len(S);i++ {
        if n,ok:=freq[S[i]];ok {
            for j:=0;j<n;j++ {
                ret = append(ret,S[i])
            }
            delete(freq,S[i])
        }
    }
    for k,n := range freq {
        for j:=0;j<n;j++ {
            ret = append(ret,k)
        }
    }
    return string(ret)
}