package main

import "fmt"
/*

Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/
// dp[n]=min(dp[k] + 1 when n == x2+k)
func numSquares(n int) int {
        dp := make([]int,n+1)
        for i:=1;i<=n;i++ {
                dp[i]=i
                for j:=1;j*j<=i;j++ {
                        if i >= j*j && i-j*j <= n && dp[i] > dp[i-j*j]+1 {
                                //加和相等 i = j*j + (i-j*j)
                                dp[i] = dp[i-j*j]+1
                        }
                }
        }
        return dp[n]
}

func main() {
        fmt.Println(numSquares(13))
}

//dp[n]=min(dp[n-x]+dp[x],x is square.)
func numSquares1(n int) int {
    dp := make([]int,n+1)
    dp[1]=1
    for i:=1;i<=n;i++ {
        dp[i]=i
        for j:=1;j*j<=i&&i-j*j<n;j++ {
            if dp[i-j*j]+1 < dp[i] {
                dp[i] = dp[i-j*j]+1
            }
        }
    }
    return dp[n]
}
