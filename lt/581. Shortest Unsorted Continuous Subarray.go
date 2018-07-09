package main

import "fmt"

/*
Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.

*/
func findUnsortedSubarray(nums []int) int {
    if len(nums) <= 0 {return 0}
    /*
    算法：看最大值和最小值哪里开始变化的。变化部分就是需要排序部分
    坑依然很多
    testcase中正序，逆序，单个元素都需要考虑在内
    */
    max := nums[0]-1
    right := len(nums)
    inorder := true
    for i:=0;i<len(nums);i++ {
        if max <= nums[i] { //一定是小于等于，否则对于相等元素情况无法通过
            max = nums[i]
        } else {
            right=i
            inorder=false
        }
    }
    if inorder {return 0}
    min := nums[len(nums)-1]
    left := -1
    for i:=len(nums)-1;i>=0;i-- {
        if min >= nums[i] {
            min=nums[i]
        } else {
            left=i
        }
    }
    //fmt.Println(nums,left,right)
    return right-left+1
}
func main() {
    fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
    fmt.Println(findUnsortedSubarray([]int{2,5,3,6}))
    fmt.Println(findUnsortedSubarray([]int{1,3,2,4,5}))
    fmt.Println(findUnsortedSubarray([]int{2,3}))
    fmt.Println(findUnsortedSubarray([]int{2,1}))
    fmt.Println(findUnsortedSubarray([]int{2}))
    fmt.Println(findUnsortedSubarray([]int{1,2,3,3,3,2}))
}
