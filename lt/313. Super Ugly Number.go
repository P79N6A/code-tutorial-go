package main

import (
    "math"
    "fmt"
)

/*

Write a program to find the nth super ugly number.

Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k.

Example:

Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12
             super ugly numbers given primes = [2,7,13,19] of size 4.
Note:

1 is a super ugly number for any given primes.
The given numbers in primes are in ascending order.
0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000.
The nth super ugly number is guaranteed to fit in a 32-bit signed integer.

*/
func nthSuperUglyNumber(n int, primes []int) int {
    if n == 0 {return 0}
    idx := make([]int,len(primes))
    ugly := make([]int,n)
    ugly[0]=1
    for i:=1;i<n;i++ {
        min := math.MaxInt64
        min_i:=0
        for j:=0;j<len(idx);j++ {
            k := primes[j]*ugly[idx[j]]
            if k < min {
                min = k
                min_i = j
            } else if k == min {
                idx[j]+=1
            }
        }
        idx[min_i]+=1
        ugly[i]=min
    }
    return ugly[n-1]
}
func main() {
    fmt.Println(nthSuperUglyNumber(0,[]int{2,7,13,19}))
    
}
