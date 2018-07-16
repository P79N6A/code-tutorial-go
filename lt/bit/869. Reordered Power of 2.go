package main

import "fmt"

/*
Starting with a positive integer N, we reorder the digits in any order (including the original order) such that the leading digit is not zero.

Return true if and only if we can do this in a way such that the resulting number is a power of 2.



Example 1:

Input: 1
Output: true
Example 2:

Input: 10
Output: false
Example 3:

Input: 16
Output: true
Example 4:

Input: 24
Output: false
Example 5:

Input: 46
Output: true


Note:

1 <= N <= 10^9
*/
func reorderedPowerOf2(N int) bool {
    /*
    题目的意思是，将十进制的数字做任意的排列，看排列之后的数字是否是2的指数
    比如 46可以转化成64 所以是true
    思路：10^9 < 2^31 结果集合较小，只有31个，看是否在结果集合中即可。数字可能会重复
    */
    for i:=0;i<=31;i++ {
        N2 := 1<<uint(i)
        if equalDigtalSlice(calDigtalSlice(N),calDigtalSlice(N2)) {
            return true
        }
    }
    return false
}
func equalDigtalSlice(a,b [10]int) bool {
    for i:=0;i<10;i++ {
        if a[i]!=b[i] {return false}
    }
    return true
}
func calDigtalSlice(N int) [10]int {
    ret := [10]int{}
    for N > 0 {
        ret[N%10] += 1
        N = N/10
    }
    return ret
}
func main() {
    fmt.Println(reorderedPowerOf2(64))
}
