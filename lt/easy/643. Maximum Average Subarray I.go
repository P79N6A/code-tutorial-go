package main

import (
    "fmt"
    "math"
)

/*
Given an array consisting of n integers, find the contiguous subarray of given length k
that has the maximum average value. And you need to output the maximum average value.

Example 1:
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
Note:
1 <= k <= n <= 30,000.
Elements of the given array will be in the range [-10,000, 10,000].
*/

func findMaxAverage(nums []int, k int) float64 {
    l,sum := 0,0
    avg := -1 *math.MaxInt64
    for i:=0;i<len(nums);i++ {
        sum += nums[i]
        if i - l == k-1 {
            if sum > avg {
                avg = sum
            }
            sum -= nums[l]
            l += 1
        }
    }
    return float64(avg)/float64(k)
}
func main() {
    fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3},4))
    fmt.Println(findMaxAverage([]int{-1},1))

}
