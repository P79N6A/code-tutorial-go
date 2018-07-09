package main

import (
    "math"
    "fmt"
)

/*
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/
// b.找连续子数组的和至少是K的最短长度，全都是正数
func minSubArrayLen(s int, nums []int) int {
    sum,left := 0,0
    ans := math.MaxInt64
    for i:=0;i<len(nums);i++ {
        sum += nums[i]
        for sum >=s {
            // 所有的sum>=s都是备选答案，ans保持备选答案中的最小值
            if ans > i-left+1 {
                ans = i-left+1
            }
            //双指针，left右移，如果不满足条件会退出sum>=s继续向前
            sum -= nums[left]
            left += 1
        }
    }
    if ans != math.MaxInt64 {return ans}
    return 0
}
func main() {
    fmt.Println(minSubArrayLen(7,[]int{2,3,1,2,4,3}))
    
}
