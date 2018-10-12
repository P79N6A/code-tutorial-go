package main

import "fmt"

/*
We have two types of tiles: a 2x1 domino shape, and an "L" tromino shape. These shapes may be rotated.

XX  <- domino

XX  <- "L" tromino
X
Given N, how many ways are there to tile a 2 x N board? Return your answer modulo 10^9 + 7.

(In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.)

Example:
Input: 3
Output: 5
Explanation:
The five different ways are listed below, different letters indicates different tiles:
XYZ XXZ XYY XXY XYY
XYZ YYZ XZZ XYY XXY
Note:

N  will be in range [1, 1000].
*/

func main() {
    fmt.Println(numTilings(2)) // 2
    fmt.Println(numTilings(3)) // 5
    fmt.Println(numTilings(4)) // 11
    fmt.Println(numTilings(5))  // 24
    fmt.Println(numTilings(6))  // 53
}
func numTilings2(N int) int {
    /*
A:dp[i][0]=dp[i-2][0]+dp[i-1][0]+dp[i-2][1]+dp[i-2][2]
B:dp[i][1]=dp[i-1][0]+dp[i-1][2]
C:dp[i][2]=dp[i-1][0]+dp[i-1][1]
    */
    if N == 1 {return 1}
    mod := int(1e9)+7
    dp := make([][]int,0)
    for i:=0;i<=N;i++ {
        dp = append(dp,make([]int,3))
    }
    dp[1][0]=1 // 长度为1, 没有凸出
    dp[1][1]=1 // 长度为1，上边凸出
    dp[1][2]=1 // 长度为1，下边凸出
    dp[2][0]=2
    dp[2][1]=2
    dp[2][2]=2
    for i:=3;i<=N;i++ {
        // 为什么不是2*dp[i-2][0] 因为这样会重复计算dp[i-1][0]部分
        dp[i][0]=(dp[i-2][0]+dp[i-1][0]+dp[i-2][1]+dp[i-2][2])%mod
        dp[i][1]=(dp[i-1][0]+dp[i-1][2])%mod
        dp[i][2]=(dp[i-1][0]+dp[i-1][1])%mod
    }
    return dp[N][0]%mod
}
/*
dp降维【为什么能够这样降维？因为BC的次序跟A中有雷同之处。】
A:dp[i][0]=(dp[i-2][0]+dp[i-1][0]+dp[i-2][1]+dp[i-2][2])%mod
B:dp[i][1]=(dp[i-1][0]+dp[i-1][2])%mod
C:dp[i][2]=(dp[i-1][0]+dp[i-1][1])%mod

将B,C带入A
dp[i][0]  =dp[i-2][0]+dp[i-1][0]+dp[i-3][0]+dp[i-3][2] + dp[i-3][0]+dp[i-3][1]
转换下顺序
             dp[i][0]  =dp[i-3][0]+dp[i-2][0]+dp[i-3][2]+dp[i-3][1] + dp[i-3][0]+dp[i-1][0]
由于i-1的dp D:dp[i-1][0]=dp[i-3][0]+dp[i-2][0]+dp[i-3][1]+dp[i-3][2]
将D带入
dp[i][0] =dp[i-1][0] + dp[i-3][0]+dp[i-1][0]
降维
dp[i]=dp[i-1]+dp[i-3]+dp[i-1]
dp[i]=2*dp[i-1]+dp[i-3]
初始条件dp[0]=0,dp[1]=1,dp[2]=2
*/
func numTilings(N int) int {
    mod := int(1e9)+7
    dp := make([]int,1000+1)
    dp[1]=1
    dp[2]=2
    dp[3]=5
    for i:=4;i<=N;i++ {
        dp[i]=(2*dp[i-1]%mod+dp[i-3])%mod
    }
    return dp[N]%mod
}