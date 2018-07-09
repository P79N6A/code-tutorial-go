package main

import "fmt"

/*
Given an array with n integers,
your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
Note: The n belongs to [1, 10,000].
*/
func checkPossibility(nums []int) bool {
    if isNonDecreasing(nums) {return true}
    for i:=1;i<len(nums);i++ {
        if nums[i-1]>nums[i] {
            a := nums[i-1]
            nums[i-1]=nums[i]
            fmt.Println(nums)
            if isNonDecreasing(nums) {
                return true
            }
            nums[i-1]= a
            a = nums[i]
            nums[i]=nums[i-1]
            fmt.Println(nums)
            if isNonDecreasing(nums) {
                return true
            }
            break
        }
    }
    return false

}
func isNonDecreasing(nums []int) bool {
    for i:=1;i<len(nums);i++ {
        if nums[i-1]>nums[i] {
            return false
        }
    }
    return true
}
func main() {
    fmt.Println(checkPossibility([]int{1,2,3}))
    //fmt.Println(checkPossibility([]int{4,2,1}))
    //fmt.Println(checkPossibility([]int{-1,4,2,3}))

}
