package main

import (
        "sort"
        "fmt"
        "math"
)
/*
Given an array nums of n integers, are there elements a, b, c in nums
such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
        sort.Ints(nums)
        res := make([][]int, 0)
        for i := 0; i < len(nums) - 2; i++ {
                l, r, sum := i + 1, len(nums) - 1, 0 - nums[i]
                if i == 0 || nums[i] != nums[i - 1] {
                        for ; l < r; {
                                if nums[l] + nums[r] == sum {
                                        res = append(res, []int{nums[i], nums[l], nums[r]})
                                        for ; l < r&&nums[l] == nums[l + 1]; {
                                                l += 1
                                        }
                                        for ; l < r&&nums[r] == nums[r - 1]; {
                                                r -= 1
                                        }
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
/*
16. 3Sum Closest

Given an array nums of n integers and an integer target,
find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

*/
func threeSumClosest(nums []int, target int) int {
        sort.Ints(nums)
        closest := 0
        min := math.MaxInt16
        for i := 0; i < len(nums) - 2; i++ {
                l, r := i + 1, len(nums) - 1
                for ; l < r; {
                        sum := nums[i] + nums[l] + nums[r]
                        diff := int(math.Abs(float64(sum-target)))
                        if diff < min {
                                min = diff
                                closest = sum
                        }
                        if sum < target {
                                l += 1
                        } else {
                                r -= 1
                        }
                }
        }
        return closest
}
func main() {
        fmt.Println(threeSum([]int{1, 2, -2, -1}))
        //fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
