package main

import (
        "fmt"
)
/*
Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.

If there is no non-empty subarray with sum at least K, return -1.



Example 1:

Input: A = [1], K = 1
Output: 1
Example 2:

Input: A = [1,2], K = 4
Output: -1
Example 3:

Input: A = [2,-1,2], K = 3
Output: 3


Note:

1 <= A.length <= 50000
-10 ^ 5 <= A[i] <= 10 ^ 5
1 <= K <= 10 ^ 9

*/
func shortestSubarray(A []int, K int) int {
        sum := make([]int,len(A)+1)
        for i:=0;i<len(A);i++ {
                sum[i+1]=A[i]+sum[i]
        }
        maxlen := -1
        for i:=0;i<len(sum);i++ {
                for j:=i+1;j<len(sum);j++ {
                        if sum[j]-sum[i] < K {
                                continue
                        }
                        if maxlen < 0 || j-i < maxlen {
                                maxlen=j-i
                        }
                }
        }
        return maxlen
}
func main() {
        fmt.Println(shortestSubarray([]int{77,19,35,10,-14},19))
        fmt.Println(shortestSubarray([]int{48,99,37,4,-31},140))
}
