package main

import (
    "math"
    "fmt"
)

/*
Given an integer array, you need to find one continuous subarray that if you only sort this subarray
in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to
make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.
*/
func findUnsortedSubarray(nums []int) int {
    left := 0
    for left+1 < len(nums) {
        if nums[left]<=nums[left+1] {
            left = left+1
        } else {
            break
        }
    }
    fmt.Println(left,len(nums))
    if left+1 >= len(nums) {return 0}
    right := len(nums)-1
    for right >= 1 {
        if nums[right-1] <= nums[right] {
            right -= 1
        } else {
            break
        }
    }
    min,max := math.MaxInt64,-1*math.MaxInt64
    for i:=left;i<=right;i++ {
        if nums[i] < min {min = nums[i]}
        if nums[i] > max {max = nums[i]}
    }
    fmt.Println(left,right,min,max)
    for left >=0 {
        if nums[left]>min {left -=1} else {break}
    }

    for right < len(nums) {
        if nums[right]<max {right+=1} else {break}
    }
    fmt.Println(left,right,min,max)
    return right-left-1
}
func findUnsortedSubarray1(nums []int) int {
    n,beg,end,min,max := len(nums),-1,-2,nums[len(nums)-1],nums[0]
    for i:=1;i<n;i++ {
        if nums[i]>max {max=nums[i]}
        if nums[n-i-1]<min {min=nums[n-1-i]}
        if nums[i]<max {end = i}
        if nums[n-1-i]>min {beg=n-1-i}
        fmt.Println(beg,end)
    }
    return end-beg+1

}
func main() {
    //fmt.Println(findUnsortedSubarray([]int{1,3,2,4,5}))
    //fmt.Println(findUnsortedSubarray([]int{1,2,3,4}))
    //fmt.Println(findUnsortedSubarray([]int{1,2,3,4,1}))
    fmt.Println(findUnsortedSubarray([]int{1,2,3,3,3}))
    //fmt.Println(findUnsortedSubarray([]int{2, 1}))
    //fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
