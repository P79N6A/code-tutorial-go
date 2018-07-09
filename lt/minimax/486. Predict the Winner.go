package main

import "fmt"

/*
Given an array of scores that are non-negative integers.
Player 1 picks one of the numbers from either end of the array followed by the player 2 and then player 1 and so on.
Each time a player picks a number, that number will not be available for the next player.
This continues until all the scores have been chosen. The player with the maximum score wins.

Given an array of scores, predict whether player 1 is the winner.
You can assume each player plays to maximize his score.

Example 1:
Input: [1, 5, 2]
Output: False
Explanation: Initially, player 1 can choose between 1 and 2.
If he chooses 2 (or 1), then player 2 can choose from 1 (or 2) and 5.
If player 2 chooses 5, then player 1 will be left with 1 (or 2).
So, final score of player 1 is 1 + 2 = 3, and player 2 is 5.
Hence, player 1 will never be the winner and you need to return False.
Example 2:
Input: [1, 5, 233, 7]
Output: True
Explanation: Player 1 first chooses 1. Then player 2 have to choose between 5 and 7.
No matter which number player 2 choose, player 1 can choose 233.
Finally, player 1 has more score (234) than player 2 (12),
so you need to return True representing player1 can win.
Note:
1 <= length of the array <= 20.
Any scores in the given array are non-negative integers and will not exceed 10,000,000.
If the scores of both players are equal, then player 1 is still the winner.
public class Solution {
    public boolean PredictTheWinner(int[] nums) {
        return winner(nums, 0, nums.length - 1, 1) >= 0;
    }
    public int winner(int[] nums, int s, int e, int turn) {
        if (s == e)
            return turn * nums[s];
        int a = turn * nums[s] + winner(nums, s + 1, e, -turn);
        int b = turn * nums[e] + winner(nums, s, e - 1, -turn);
        return turn * Math.max(turn * a, turn * b);
    }
}
*/

func PredictTheWinner(nums []int) bool {
    solve(nums,0,len(nums)-1,1)
    solve3(nums)
    return solve2(nums,0,len(nums)-1)>=0
}
func solve(nums []int, s,e int, turn int) int {
    if s == e {
        fmt.Println(turn*nums[s])
        return  turn *nums[s]
    }
    a := turn * nums[s] + solve(nums,s+1,e,-turn)
    b := turn * nums[e] + solve(nums,s,e-1,-turn)
    if turn*a > turn*b {
        return a
    }
    return b
}
//谁都去得是当前的最大收益
func solve2(nums []int, s,e int) int {
    if s == e {
        return  nums[s]
    }
    a := nums[s] - solve2(nums,s+1,e)
    b := nums[e] - solve2(nums,s,e-1)
    if a > b {
        fmt.Println(a,nums[s])
        return a
    }
    fmt.Println(b,nums[e])
    return b
}
/*
        int[][] dp = new int[nums.length + 1][nums.length];
        for (int s = nums.length; s >= 0; s--) {
            for (int e = s + 1; e < nums.length; e++) {
                int a = nums[s] - dp[s + 1][e];
                int b = nums[e] - dp[s][e - 1];
                dp[s][e] = Math.max(a, b);
            }
        }
        return dp[0][nums.length - 1] >= 0;
*/
func solve3(nums []int) {
    dp := make([][]int,0)
    for i:=0;i<len(nums)+1;i++ {
        dp = append(dp,make([]int,len(nums)))
    }
    for s:=len(nums);s>=0;s-- {
        for e := s+1;e < len(nums);e++ {
            a := nums[s]-dp[s+1][e]
            b := nums[e]-dp[s][e-1]
            if a > b {
                dp[s][e] = a
            } else {
                dp[s][e] = b
            }
        }
    }
    for _,d := range dp {
        fmt.Println(d)
    }

}


func main() {
    //fmt.Println(PredictTheWinner([]int{1,5,8,4}))
    fmt.Println(PredictTheWinner([]int{2,4,6}))

}