package main

import "fmt"

func longestPalindromeSubseq(s string) int {
        slen := len(s)
        dp := make([][]int,slen)
        for i:=0;i<slen;i++ {
                dp[i]=make([]int,slen)
        }
        /*
         右上三角形;原因在于状态转换方程
        */
        for i:=slen-1;i>=0;i-- {
                dp[i][i] = 1;
                for j:=i+1;j<slen;j++ {
                        if s[i]==s[j] {
                                dp[i][j] = dp[i+1][j-1] + 2
                        } else {
                                if dp[i+1][j] > dp[i][j-1]{
                                        dp[i][j]=dp[i+1][j]
                                } else {
                                        dp[i][j]=dp[i][j-1]
                                }
                        }
                }
        }
        return dp[0][slen-1]

}

func longestPalindromeSubseq2(s string) int {
        slen := len(s)
        dp := make([][]int,slen)
        for i:=0;i<slen;i++ {
                dp[i]=make([]int,slen)
        }
        /*
         左上三角形;原因在于状态转换方程

         i-j范围是从0-slen,所以是从中间的对角线开始计算的
         关键在于区间类型动态规划的状态方程,i和j的关系映射到二维数据上
        */
        for i:=0;i<slen;i++ {
                dp[i][i] = 1;
                for j:=i-1;j>=0;j-- {
                        if s[i]==s[j] {
                                dp[i][j] = dp[i-1][j+1] + 2
                        } else {
                                if dp[i-1][j] > dp[i][j+1]{
                                        dp[i][j]=dp[i-1][j]
                                } else {
                                        dp[i][j]=dp[i][j+1]
                                }
                        }
                }
        }
        return dp[slen-1][0]

}
func main() {
        fmt.Println(longestPalindromeSubseq2("cbbd"))
}
