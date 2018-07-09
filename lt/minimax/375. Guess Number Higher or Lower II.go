package main

import (
    "fmt"
    "math"
)

/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number I picked is higher or lower.

However, when you guess a particular number x, and you guess wrong, you pay $x. You win the game when you guess the number I picked.

Example:

n = 10, I pick 8.

First round:  You guess 5, I tell you that it's higher. You pay $5.
Second round: You guess 7, I tell you that it's higher. You pay $7.
Third round:  You guess 9, I tell you that it's lower. You pay $9.

Game over. 8 is the number I picked.

You end up paying $5 + $7 + $9 = $21.
Given a particular n ≥ 1, find out how much money you need to have to guarantee a win.

Credits:
Special thanks to @agave and @StefanPochmann for adding this problem and creating all test cases.
*/

/*
还是穷举所有情况。也就是我们假设结果可能是任意一个数字
因为每次猜完都会将数组分成两部分，需要付出的代价是两部分的较大者。
那就看这次猜那个数字能将这个较大者变的最小
*/
func getMoneyAmount(n int) int {
    return solve(1,n,make(map[string]int))
}
func solve(left,right int,cache map[string]int) int {
    // 从left-right 猜中数字最多需要花费多少。计算所有可能情况。
    // 如果只有一个数字，则不需要花费，直接猜对
    // 如果两个数字，则肯定是先猜较小那个，因为即使猜错了下次也知道正确答案了
    if left >= right {return 0}
    key := fmt.Sprintf("%d-%d",left,right)
    if _,ok := cache[key];ok {return cache[key]}
    ret := math.MaxInt64
    for i:=left;i<=right;i++ {
        // 猜每一种可能i，的最坏情况中的最好情况
        m := i + max(solve(left,i-1,cache),solve(i+1,right,cache)) // 最坏情况，因为要解决就得处理得了最坏的
        ret = min(ret,m) // 留存最好情况，猜i
    }
    cache[key]=ret
    return ret
}
func max(a,b int) int {
    if a > b {return a}
    return b
}
func min(a,b int) int {
    if a > b {return b}
    return a
}
func main() {
    for i:=0;i<100;i++ {
        fmt.Println(i,getMoneyAmount(i))
    }

}
