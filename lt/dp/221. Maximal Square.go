package main

import "fmt"
/*
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4

*/

func main() {
    fmt.Println(maximalSquare([][]byte{
        []byte("10100"),
        []byte("10111"),
        []byte("11111"),
        []byte("10010"),
    }))
}
func maximalSquare(matrix [][]byte) int {
    if len(matrix)<=0 {return 0}
    dp := make([][]int,0)
    for i:=0;i<=len(matrix);i++{
        dp = append(dp,make([]int,len(matrix[0])+1))
    }
    ans := 0
    for i:=1;i<=len(matrix);i++ {
        for j:=1;j<=len(matrix[i-1]);j++ {
            if matrix[i-1][j-1]=='1' {
                // 正方形,长宽要相等
                // dp[i][j] 代表以i,j为正方形的右下顶点的最大边长
                dp[i][j]=min(dp[i-1][j],min(dp[i-1][j-1],dp[i][j-1]))+1
                ans = max(ans,dp[i][j])
            }
        }
    }
    return ans*ans
}
func min(a,b int) int {
    if a<b {return a}
    return b
}
func max(a,b int) int {
    if a<b {return b}
    return a
}
