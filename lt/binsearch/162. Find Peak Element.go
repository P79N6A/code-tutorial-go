package main

import "fmt"

/*

A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 1 or 5
Explanation: Your function can return either index number 1 where the peak element is 2,
             or index number 5 where the peak element is 6.
*/
func findPeakElement(nums []int) int {
    if len(nums)<=1 {return 0}
    l,r := 0,len(nums)
    for l<r {
        m := (l+r)/2
        if m == 0 {
            if nums[m]>nums[m+1] {
                return m
            } else {
                l=m+1
            }
        } else if m == len(nums)-1 {
            if nums[m]>nums[m-1] {
                return m
            } else {
                r=m
            }
        } else {
            if nums[m] > nums[m-1] && nums[m]>nums[m+1] {
                return m
            } else if nums[m] > nums[m-1] && nums[m]<nums[m+1] {
                l = m+1
            } else {
                r=m
            }
        }
    }
    return r
}
func main() {
    fmt.Println(findPeakElement([]int{1,2,1,3,5,6,4}))
    fmt.Println(findPeakElement([]int{1,2,3,1}))
    fmt.Println(findPeakElement([]int{1,2,3}))
    fmt.Println(findPeakElement([]int{5,4,3}))

}
