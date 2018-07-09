package main

import (
    "fmt"
    "math"
)

/*
Given an array which consists of non-negative integers and an integer m,
you can split the array into m non-empty continuous subarrays.
Write an algorithm to minimize the largest sum among these m subarrays.

Note:
If n is the length of array, assume the following constraints are satisfied:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
Examples:

Input:
nums = [7,2,5,10,8]
m = 2

Output:
18

Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.

*/
// dp[m][n]=min(max(dp[m-1][k],sum[n-k])) k start at m-1=>n
// k start m-1 means : make sure m-1 must have m-1 element.
// dp[m][n] not dp[n][m] 原因是 每次都是m + 1 做计算
// 这个可以做数组压缩


func splitArray(nums []int, m int) int {
    dp := make([][]int, 0)
    sum := make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        sum[i] = sum[i-1] + nums[i-1]
    }
    for i := 0; i <= m; i++ {
        dp = append(dp, make([]int, len(nums)+1))
    }
    for i := 0; i <= m; i++ {
        for j := 0; j <= len(nums); j++ {
            dp[i][j] = math.MaxInt64
        }
    }
    dp[0][0] = 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= len(nums); j++ {
            for k := i - 1; k < j; k++ {
                //for k:=j-1;k>=0;k-- {
                max := dp[i-1][k]
                if max < sum[j]-sum[k] {
                    max = sum[j] - sum[k]
                }
                if dp[i][j] > max {
                    dp[i][j] = max
                }
            }
        }
        for _,d := range dp {
            fmt.Println(d)
        }
        fmt.Println("XXXXXXX")
    }
    return dp[m][len(nums)]
}
func main() {
    //fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
    fmt.Println(splitArray([]int{1, 2,2147483646}, 2))

}
