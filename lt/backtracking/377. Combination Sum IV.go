package main

import (
    "fmt"
    "sort"
)

/*
Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.

Example:

nums = [1, 2, 3]
target = 4

The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

Note that different sequences are counted as different combinations.

Therefore the output is 7.
Follow up:
What if negative numbers are allowed in the given array?
How does it change the problem?
What limitation we need to add to the question to allow negative numbers?

Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.

dp[i]=sum(dp[i-x],x in nums)
*/

func combinationSum4(nums []int, target int) int {
        if len(nums) <= 0 {return 0}
        dp := make([]int,target+1)
        sort.Ints(nums)
        dp[0]=1
        for i:=1;i<=target;i++ {
                for _,n := range nums {
                        if i-n>=0{
                                dp[i]+=dp[i-n]
                        }
                }
        }
        fmt.Println(dp)
        return dp[target]
}
func combinationSum41(nums []int, target int) int {
        num := 0
        ret := make([]int,0)
        solve(nums,0,target,&num,&ret)
        return num
}
func solve(nums []int, j int, target int,num *int,ret *[]int)   {
        if target == 0 {
                *num += 1
                fmt.Println(*ret)
                return
        }
        if target < 0 {
                return
        }
        for i:=j;i<len(nums);i++ {
                *ret = append(*ret,nums[i])
                solve(nums, j, target - nums[i], num,ret)
                *ret = (*ret)[:len(*ret)-1]
                solve(nums, j + 1, target, num,ret)
        }
}

func combinationSum4(nums []int, target int) int {
    dp := make([]int,target+1)
    dp[0]=1
    for i:=0;i<len(dp);i++ {
        for j:=0;j<len(nums);j++ {
            if i >= nums[j] {
                dp[i] += dp[i-nums[j]]
            }
        }
    }
    return dp[target]
}
func combinationSum44(nums []int, target int) int {
    if len(nums) <= 0 {return 0}
    dp := make([]int,target+1)
    sort.Ints(nums)
    dp[0]=1
    for i:=1;i<=target;i++ {
        for _,n := range nums {
            if i-n>=0{
                dp[i]+=dp[i-n]
            }
        }
    }
    fmt.Println(dp)
    return dp[target]
}
func combinationSum41(nums []int, target int) int {
    ret := make([][]int,0)
    res := make([]int,0)
    bt(nums,target,0,&res,&ret)
    fmt.Println(ret)
    return len(ret)
}
func bt(nums []int,target int,sum int,res *[]int, ret *[][]int) {
    if sum > target {return}
    if sum == target {
        fmt.Println(res)
        n := make([]int,len(*res))
        copy(n,*res)
        *ret = append(*ret,n)
        return
    }
    for i:=0;i<len(nums);i++ {
        *res = append(*res,nums[i])
        bt(nums,target,sum+nums[i],res,ret)
        *res = (*res)[:len(*res)-1]
    }
}
func main() {
    fmt.Println(combinationSum4([]int{4,2,1},32))
    fmt.Println(combinationSum44([]int{4,2,1},32))
}
