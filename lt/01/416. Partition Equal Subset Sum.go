package main

import "fmt"
/*
Given a non-empty array containing only positive integers,
find if the array can be partitioned into two subsets
such that the sum of elements in both subsets is equal.

Note:
Each of the array element will not exceed 100.
The array size will not exceed 200.
Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.

*/
func canPartition(nums []int) bool {
        sum := 0
        for _,n := range nums {sum += n}
        if sum % 2 == 1 {
                return false
        }
        w := sum / 2
        dp := make([]bool, w + 1)
        dp[0] = true
        // dp[i][w]=dp[i-1][w] || dp[i][w-nums[i]]
        for _, num := range nums {
                for i := w; i >= 0; i-- {
                        if i >= num {
                                dp[i] = dp[i] || dp[i - num]
                        }
                }
        }
        return dp[w]
}
func main() {
        //fmt.Println(canPartition([]int{11,5,11,5}))
        fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
// dp[i][sum]=dp[i-1][sum] || dp[i-1][sum-num[i]]
func canPartition1(nums []int) bool {
        sum := 0
        for _, n := range nums {
                sum += n
        }
        if sum % 2 != 0 {
                return false
        }
        sum = sum / 2
        dp := make([]bool, sum + 1)
        dp[0] = true
        for _, n := range nums {
                for j := sum; j >= 0 && j >= n; j-- {
                        dp[j] = dp[j] || dp[j - n]
                }
        }
        return dp[sum]
}

