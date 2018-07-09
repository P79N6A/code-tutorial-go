package main

import (
        "math"
        "fmt"
)
/*
Given two words word1 and word2,
find the minimum number of steps required to make word1 and word2 the same,
where in each step you can delete one character in either string.

Example 1:
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Note:
The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.
*/
// dp[i][j] = min(dp[i][j-1] + 1, dp[i-1][j] +1) if dp[i] != dp[j]
// dp[i][j] = dp[i-1][j-1] if dp[i]==dp[j]
/*
编辑距离和最长公共子序列相同的问题
*/
func minDistance(word1 string, word2 string) int {
        w1len,w2len := len(word1),len(word2)
        dp := make([][]int,w1len+1)
        for i:=0;i<w1len+1;i++ {
                dp[i]=make([]int,w2len+1)
                dp[i][0]=i
        }
        //边界初始化,因为是min方法,所以边界不能是0,应该是至少需要的步数,也就是index的值.
        for j:=0;j<w2len+1;j++ {dp[0][j]=j}
        for i:=1;i<=w1len;i++ {
                for j:=1;j<=w2len;j++ {
                        if word1[i-1] == word2[j-1] {
                                dp[i][j]=dp[i-1][j-1]
                        } else {
                                fmt.Println(i,j,w1len,w2len)
                                dp[i][j]=int(math.Min(float64(dp[i][j-1]+1),float64(dp[i-1][j]+1)))
                        }
                }
        }
        return dp[w1len][w2len]
}

func main() {
        fmt.Println(minDistance("sea","eat"))
}
