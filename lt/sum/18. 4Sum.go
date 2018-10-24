package main

import (
    "sort"
    "fmt"
)

/*
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func main() {
    fmt.Println(fourSum([]int{1,0,-1,0,-2,2},0))
    
}
func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    ret := make([][]int,0)
    for i:=0;i<len(nums)-3;i++ {
        if i == 0 || nums[i]!=nums[i-1] {
            x := threeSum(nums[i+1:],target-nums[i])
            for _,r := range x {
                r = append(r,nums[i])
                sort.Ints(r)
                ret = append(ret,r)
            }
        }
    }
    return ret
}

func threeSum(nums []int,target int) [][]int {
    res := make([][]int,0)
    for i := 0;i < len(nums)-2;i++ {
        l,r,sum := i+1,len(nums)-1,target-nums[i]
        if i == 0 || nums[i] != nums[i-1] {
            for ; l < r; {
                if nums[l] + nums[r] == sum {
                    res = append(res, []int{nums[i], nums[l], nums[r]})
                    for ;l<r&&nums[l]==nums[l+1];{l += 1}
                    for ;l<r&&nums[r]==nums[r-1];{r -= 1}
                    l += 1
                    r -= 1
                } else if nums[l] + nums[r] < sum {
                    l += 1
                } else {
                    r -= 1
                }
            }
        }
    }
    return res
}