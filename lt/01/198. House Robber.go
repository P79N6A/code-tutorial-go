package main

import "fmt"

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is
that adjacent houses have security system connected and
it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

// dp[i]=max(dp[i-1],dp[i-2]+num[i])
func rob(nums []int) int {
    /*
    背包问题在空间上都可以降级的，降低一个维度。
    */
    a,b,dp := 0,0,0
    for i:=0;i<len(nums);i++ {
        if b < a + nums[i] {
            dp = a + nums[i]
        } else {
            dp = b
        }
        a,b = b,dp
    }
    return dp
}

func main() {
    fmt.Println(rob([]int{1}))
    fmt.Println(rob([]int{1,2,3,1}))
    fmt.Println(rob([]int{2,1,1,2}))
    fmt.Println(rob([]int{2,7,9,3,1}))

}
