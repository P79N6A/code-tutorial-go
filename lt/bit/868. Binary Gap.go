package main

import (
    "fmt"
    "math"
)

/*
Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.



Example 1:

Input: 22
Output: 2
Explanation:
22 in binary is 0b10110.
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
The first consecutive pair of 1's have distance 2.
The second consecutive pair of 1's have distance 1.
The answer is the largest of these two distances, which is 2.
Example 2:

Input: 5
Output: 2
Explanation:
5 in binary is 0b101.
Example 3:

Input: 6
Output: 1
Explanation:
6 in binary is 0b110.
Example 4:

Input: 8
Output: 0
Explanation:
8 in binary is 0b1000.
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.
*/
// 需要注意的是，需要至少两个1
func binaryGap(N int) int {
    mcon,con := 0,0
    for N>0&&N%2==0 {N = N>>1}
    if N ==1 {return 0}
    for N > 0 {
        if N%2 ==0  {
            con += 1
        } else {
            if con > mcon {
                mcon = con
            }
            con = 0
        }
        N = N>>1
    }
    if con > mcon {
        mcon = con
    }
    return mcon+1
}
func main() {
    var ii int64 = math.MinInt32
    fmt.Println(ii)
    fmt.Println(math.Abs(float64(ii)))
    fmt.Println(-1*ii)
    fmt.Println(binaryGap(8))
}
