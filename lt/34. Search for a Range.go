package main

import "fmt"
/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
func binarySearch(nums []int,target int) int {
        l,r := 0,len(nums)
        for ;l<r; {
                // 有可能越界
                m := l + (l-r)/2
                if nums[m] == target {
                        return m
                }else if nums[m] <= target {
                        l = m + 1
                } else {
                        r = m
                }
        }
        return -1
}
func upperBound(nums []int,target int) int {
        l,r := 0,len(nums)
        for ;l<r; {
                m := (l+r)/2
                // <= 意味着,即使等于target,l也会继续右移,那么r就是第一个大于target的下标
                if nums[m] <= target {
                        l = m + 1
                } else {
                        r = m
                }
        }
        return r
}
func lowerBound(nums []int,target int) int {
        l,r := 0,len(nums)
        for ;l<r; {
                m := (l+r)/2
                //和upperbound唯一的区别
                // < 意味着,l会右移到小于target,那么r就是第一个等于target的下标,如果没有等于就是大于
                if nums[m] < target {
                        l = m + 1
                } else {
                        r = m
                }
        }
        return r
}

func binsearch_upper(nums []int,t int) int {
        l,r := 0,len(nums)
        for ;l<r; {
                m := (l+r)/2
                if nums[m] < t {
                        l = m + 1
                } else {
                        r = m
                }
        }
        // r === l
        fmt.Println("yyy",l,r)
        // r is the first not bigger then t
        if r < len(nums) && nums[r] == t {
                return r
        }
        return -1
}
func binsearch_lower(nums []int,t int) int {
        l,r := 0,len(nums)
        for {
                if l >= r {
                        break
                }
                m := (l+r)/2
                if nums[m] <= t {
                        l = m + 1
                } else {
                        r = m
                }
        }
        // r === l
        fmt.Println("xxx",l,r)
        // r is the first bigger then t
        if r > 0 && nums[r-1] == t {
                return r-1
        }
        return -1
}

func searchRange(nums []int, target int) []int {
        a := binsearch_upper(nums,target)
        b := binsearch_lower(nums,target)
        return []int{a,b}
}

func main() {
        fmt.Println(searchRange([]int{5,7,7,7,8,10},8))
        //fmt.Println(searchRange(nil,0))
        fmt.Println(upperBound([]int{1,2,3,3,4},4))
        fmt.Println(lowerBound([]int{1,2,3,3,4},4))
}
