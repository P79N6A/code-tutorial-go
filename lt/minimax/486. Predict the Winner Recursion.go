package main

import "fmt"

/*
Given an array of scores that are non-negative integers. Player 1 picks one of the numbers from either end of the array followed by the player 2 and then player 1 and so on. Each time a player picks a number, that number will not be available for the next player. This continues until all the scores have been chosen. The player with the maximum score wins.

Given an array of scores, predict whether player 1 is the winner. You can assume each player plays to maximize his score.

Example 1:
Input: [1, 5, 2]
Output: False
Explanation: Initially, player 1 can choose between 1 and 2.
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5. If player 2 chooses 5, then player 1 will be left with 1 (or 2).
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5.
Hence, player 1 will never be the winner and you need to return False.
Example 2:
Input: [1, 5, 233, 7]
Output: True
Explanation: Player 1 first chooses 1. Then player 2 have to choose between 5 and 7. No matter which number player 2 choose, player 1 can choose 233.
Finally, player 1 has more score (234) than player 2 (12), so you need to return True representing player1 can win.
Note:
1 <= length of the array <= 20.
Any scores in the given array are non-negative integers and will not exceed 10,000,000.
If the scores of both players are equal, then player 1 is still the winner.

*/
func PredictTheWinner(nums []int) bool {
    return canWin(1,0,0,nums)
}
func canWin(player int,sumA,sumB int,nums[]int) bool {
    // 函数返回当前的player能否赢
    if len(nums) <= 0 {
        //退出的情况 如果是player1，则sumA>=sumB才算赢
        if player == 1 {
            return sumA >= sumB
        }
        return sumB > sumA
    }
    if len(nums) == 1 {
        if player == 1 {
            return sumA+nums[0]>=sumB
        }
        //退出情况，还剩一个元素， 如果是player2，那么sumB把最后数字取了，然后大于sumA才算赢。
        return sumB + nums[0] > sumA
    }
    if player == 1 {
        // player1能赢的两种情况，第一种是player2在player1取了左侧的时候输了
        // 第二种情况是player2在player1取了右侧的时候输了
        if !canWin(2,sumA+nums[0],sumB,nums[1:]) {
            return true
        }
        if !canWin(2,sumA+nums[len(nums)-1],sumB,nums[:len(nums)-1]) {
            return true
        }
    } else {
        if !canWin(1,sumA,sumB+nums[0],nums[1:]) {
            return true
        }
        if !canWin(1,sumA,sumB+nums[len(nums)-1],nums[:len(nums)-1]) {
            return true
        }
    }
    return false
}


func main() {
    fmt.Println(PredictTheWinner([]int{1, 1}))//true
    fmt.Println(PredictTheWinner([]int{1, 5, 2}))
    fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))

}
