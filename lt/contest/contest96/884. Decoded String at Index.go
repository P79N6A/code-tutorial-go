package main

import "fmt"

/*
An encoded string S is given.  To find and write the decoded string to a tape, the encoded string is read one character at a time and the following steps are taken:

If the character read is a letter, that letter is written onto the tape.
If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.



Example 1:

Input: S = "leet2code3", K = 10
Output: "o"
Explanation:
The decoded string is "leetleetcodeleetleetcodeleetleetcode".
The 10th letter in the string is "o".
Example 2:

Input: S = "ha22", K = 5
Output: "h"
Explanation:
The decoded string is "hahahaha".  The 5th letter is "h".
Example 3:

Input: S = "a2345678999999999999999", K = 1
Output: "a"
Explanation:
The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".


Note:

2 <= S.length <= 100
S will only contain lowercase letters and digits 2 through 9.
S starts with a letter.
1 <= K <= 10^9
The decoded string is guaranteed to have less than 2^63 letters.
*/

func main() {
    fmt.Println(decodeAtIndex("ha22", 5))
    fmt.Println(decodeAtIndex("leet2code3", 10))
    fmt.Println(decodeAtIndex("a2345678999999999999999", 10))
    fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 10))
    fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 3))
    fmt.Println(decodeAtIndex("a2b3c4d5e6f7g8h9", 9)) //b
    fmt.Println(decodeAtIndex("a2b", 3))
    //        fmt.Println(decodeLen("leet2code3",lenCache))
}
var lenCache = make(map[string]int)
func decodeAtIndex(S string, K int) string {
    i := len(S) - 1
    for ; i >= 0; i-- {
        // find alpha
        if !(S[i] >= '0' && S[i] <= '9') {break}
    }
    K = K % decodeLen(S[:i+1], lenCache)
    if K == 0 {
        // so many bugs...
        // 取余在最后一个数字的时候是0
        //  a2c 3 ;    abcd 4
        return string(S[i])
    }
    n := i
    for ;n >= 0; n-- {
        // find number
        if (S[n] >= '0' && S[n] <= '9') {
            break
        }
    }
    sublen := decodeLen(S[:n+1], lenCache)
    if K > sublen {
        for j := n; j < len(S); j++ {
            if sublen+j-n == K {return string(S[j])}
        }
    }
    return decodeAtIndex(S[:n], K)
}
func decodeLen(S string, cache map[string]int) int {
    if len(S) <= 0 {
        return 0
    }
    if cache[S] > 0 {
        return cache[S]
    }
    subfixAlpha := 0
    i := 0
    for i = len(S) - 1; i >= 0; i-- {
        if S[i] >= '0' && S[i] <= '9' {
            break
        } else {
            subfixAlpha += 1
        }
    }
    ret := subfixAlpha
    if i > 0 {
        ret = subfixAlpha + int(S[i]-'0')*decodeLen(S[:i], cache)
    }
    cache[S] = ret
    return ret
}
