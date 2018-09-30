package main

import "fmt"

/*
You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.

Note: You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer


Example 1:

Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1


Example 2:

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.


Example 3:

Input: amount = 10, coins = [10]
Output: 1

*/

func main() {
    fmt.Println(change(5,[]int{1,2,5}))
    
}
// 这是个完全背包问题，coin是物件；amount是容量，物件一个一个放进来。逐步的更新容量
// 普通的递归无法去重复，这类递归保证每次处理只获取一个最新的coin，从小到大是否能达到指定amount
// 二维的dp还是需要理解下。
func change(amount int, coins []int) int {
    dp := make([][]int,0)
    for i:=0;i<=len(coins);i++ {
        dp = append(dp,make([]int,amount+1))
    }
    dp[0][0]=1
    // 硬币一个一个加入，所以在最外层循环
    for i:=1;i<=len(coins);i++ {
        dp[i][0]=1
        //随着硬币加入有两种情况
        /*
        dp[i][j]代表前i个硬币 达到j的数值，有多少种组合方法
        只包含一个新加入到硬币；并处理j增加到amount的情况=》这个case可以保证是唯一的
        当j>当前硬币 dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
        否则，就是当前硬币加入不进来 dp[i][j] = dp[i-1][j]
        */
        for j:=1;j<=amount;j++ {
            if j>=coins[i-1] {
                dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
            } else {
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[len(coins)][amount]
}

