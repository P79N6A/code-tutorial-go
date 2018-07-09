package main

import "fmt"

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?

*/
func rotate(nums []int, k int)  {
    k = k%len(nums)
    rotateArry(nums,0,len(nums))
    rotateArry(nums,0,k)
    rotateArry(nums,k,len(nums))
}
func rotateArry(nums []int,start,end int) {
    fmt.Println(start,end,len(nums))
    for start < end-1 && len(nums) >= end {
        nums[start],nums[end-1]=nums[end-1],nums[start]
        start,end = start+1,end-1
    }
}
func main() {
    nums := []int{1,2,3}
    //nums := []int{1}
    //nums := []int{1,2,3,4,5,6,7}
    rotate(nums,4)
    fmt.Println(nums)

}
