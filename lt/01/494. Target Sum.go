package main

import "fmt"
/*
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S.
Now you have 2 symbols + and -.
For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
Note:
The length of the given array is positive and will not exceed 20.
The sum of elements in the given array will not exceed 1000.
Your output answer is guaranteed to be fitted in a 32-bit integer.

A-B=target
A+B=sum
A=(target+sum)/2
*/
func findTargetSumWays(nums []int, S int) int {
        sum := 0
        for i:=0;i<len(nums);i++ {
                sum += nums[i]
        }
        if S>sum || (sum+S)%2 == 1 {
                return 0
        }
        w := (S+sum)/2
        dp := make([]int,w+1)
        dp[0]=1
        // dp[i][w] = dp[i-1][w] + dp[i-1][w-nums[i]]
        // i 可以去掉,只有一个数组就可以了,保证从后往前计算.
        for i:=0;i<len(nums);i++ {
                for j:=w;j>=0;j-- {
                        if j >= nums[i] {
                                dp[j]=dp[j]+dp[j-nums[i]]
                        }
                }
        }
        return dp[w]

}
func main() {
        //fmt.Println(findTargetSumWays([]int{1,1,1,1,1},3))
        fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
/*
A-B = t
A+B-2B = t
sum - 2B = t
B = (sum-t)/2
转换：从nums中选出若干个数，满足和是B
dp[i][B] 表示从i个数总选择，满足和是B的所有方式
dp[i][B]=dp[i-1][B] + dp[i-1][B-nums[i]]  // 选择i或者不选择i的方法加在一起
dp[j]=dp[j]+dp[j-nums[i]]
*/
func findTargetSumWays1(nums []int, S int) int {
    sum := 0
    for _,n:=range nums {sum += n}
    if (sum - S) < 0 || (sum-S)%2 == 1 {
        return 0
    }
    w := (sum-S)/2
    dp := make([]int,w+1)
    dp[0]=1
    for _,n := range nums {
        for j := w;j>=0&&j>=n;j-- {
            dp[j] = dp[j]+dp[j-n]
        }
    }
    return dp[w]
}
