package main

import (
        "fmt"
        "strconv"
        "math"
)
/*
Given a string S of digits, such as S = "123456579", we can split it into a Fibonacci-like sequence [123, 456, 579].

Formally, a Fibonacci-like sequence is a list F of non-negative integers such that:

0 <= F[i] <= 2^31 - 1, (that is, each integer fits a 32-bit signed integer type);
F.length >= 3;
and F[i] + F[i+1] = F[i+2] for all 0 <= i < F.length - 2.
Also, note that when splitting the string into pieces, each piece must not have extra leading zeroes, except if the piece is the number 0 itself.

Return any Fibonacci-like sequence split from S, or return [] if it cannot be done.

Example 1:

Input: "123456579"
Output: [123,456,579]
Example 2:

Input: "11235813"
Output: [1,1,2,3,5,8,13]
Example 3:

Input: "112358130"
Output: []
Explanation: The task is impossible.
Example 4:

Input: "0123"
Output: []
Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.
Example 5:

Input: "1101111"
Output: [110, 1, 111]
Explanation: The output [11, 0, 11, 11] would also be accepted.
Note:

1 <= S.length <= 200
S contains only digits.
*/

func splitIntoFibonacci(S string) []int {
        for i:=1;i<len(S);i++ {
                for j:=i+1;j<len(S)-1;j++ {
                        p,_ := strconv.Atoi(S[:i])
                        q,_ := strconv.Atoi(S[i:j])
                        if p > math.MaxInt32 || q > math.MaxInt32 {
                                continue
                        }
                        ok,ret := split(p,q,S[j:])
                        if ok {
                                result := []int{p,q}
                                result = append(result,ret...)
                                sl := 0
                                for _,r:= range result {
                                        rs := strconv.Itoa(r)
                                        sl += len(rs)
                                }
                                if sl != len(S) {
                                        return nil
                                }
                                return result
                        }
                }
        }
        return nil
}
func split(one,two int, S string) (bool,[]int) {
        if len(S) <= 0 {
                return true,nil
        }
        result := make([]int,0)
        si,_ := strconv.Atoi(S)
        if si < two {return false,nil}
        for i:=1;i<len(S);i++ {
                si,_ := strconv.Atoi(S[:i])
                if si > math.MaxInt32 {
                        continue
                }
                if one + two == si {
                        ok,less := split(two,si,S[i:])
                        if ok {
                                result = append(result,si)
                                result = append(result,less...)
                                return true,result
                        }
                }
        }
        if one+two == si && si < math.MaxInt32 {return true,[]int{si}}
        return false,result
}

func main() {
        //fmt.Println(splitIntoFibonacci("123456579"))
        //fmt.Println(splitIntoFibonacci("020406"))
        //fmt.Println(splitIntoFibonacci("0000"))
        fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
        fmt.Println(splitIntoFibonacci("01123581321345589"))
        fmt.Println(math.MaxInt32)
}
