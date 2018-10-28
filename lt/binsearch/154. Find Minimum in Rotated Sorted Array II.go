package main

import "fmt"
/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:

Input: [1,3,5]
Output: 1
Example 2:

Input: [2,2,2,0,1]
Output: 0
Note:

This is a follow up problem to Find Minimum in Rotated Sorted Array.
Would allow duplicates affect the run-time complexity? How and why?
*/

func main() {
        fmt.Println(findMin([]int{2,2,2,0,1}))
        fmt.Println(findMin([]int{1,3,5}))
}
func min(a,b int)int {
        if a<b {return a}
        return b
}
func findMin(nums []int) int {
        if len(nums)==1{return nums[0]}
        if nums[0]<nums[len(nums)-1]{return nums[0]} // 已经有序
        if nums[0]==nums[len(nums)-1] { //无法二分，那就分治到子问题
                m := len(nums)/2
                x := findMin(nums[:m])
                y := findMin(nums[m:])
                return min(x,y)
        }
        m := len(nums)/2
        if m>0&&nums[m]<nums[m-1]{ // 找到最小值
                return nums[m]
        }
        if m<len(nums)-1&& nums[m]>nums[m+1] { // 找到最大值，最大旁边就是最小
                return nums[m+1]
        }
        if nums[m]>=nums[0] {
                return findMin(nums[m:])
        }
        return findMin(nums[:m])
}

