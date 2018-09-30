package main

import (
    "math"
    "fmt"
)

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.


*/

func main() {
    fmt.Println(coinChange([]int{1,2,5},11))
    fmt.Println(coinChange([]int{2},3))
}
func coinChange(coins []int, amount int) int {
    // 区间dp
    // dp[i]=min(dp[i-x]+1) x in coins
    dp := make([]int,amount+1)
    for i:=1;i<=amount;i++ {
        dp[i]=math.MaxInt64
        for _,c := range coins {
            if i-c>=0{
                dp[i]=min(dp[i],dp[i-c])
            }
        }
        if dp[i]!=math.MaxInt64 {dp[i]+=1} // valid
    }
    if dp[amount]==math.MaxInt64{return -1}
    return dp[amount]
}
func min(a,b int) int {
    if a>b {return b}
    return a
}
func coinChangeR(coins []int, amount int) int {
    ok,m := solve(coins,amount,make(map[int]int))
    if !ok {return -1}
    return m
}
func solve(coins []int, amount int,cache map[int]int) (bool,int) {
    if amount == 0 {return true,0}
    if _,ok := cache[amount];ok{
        if cache[amount]==math.MaxInt64 {
            return false,-1
        }
        return true,cache[amount]
    }
    min := math.MaxInt64
    for i:=0;i<len(coins);i++ {
        if amount >= coins[i] {
            ok,x := solve(coins,amount-coins[i],cache)
            if ok&&x < min {min = x}
        }
    }
    if min == math.MaxInt64 {
        cache[amount]=min
        return false,-1
    }
    cache[amount]=min+1
    return true,min+1
}