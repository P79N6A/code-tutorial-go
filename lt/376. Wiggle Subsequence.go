package main

import "fmt"

/*

A sequence of numbers is called a wiggle sequence if the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

Examples:
Input: [1,7,4,9,2,5]
Output: 6
The entire sequence is a wiggle sequence.

Input: [1,17,5,10,13,15,10,5,16,8]
Output: 7
There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].

Input: [1,2,3,4,5,6,7,8,9]
Output: 2
Follow up:
Can you do it in O(n) time?
*/
func wiggleMaxLength(nums []int) int {
    if len(nums) <= 0 {return 0}
    up,down := 1,1
    for i:=0;i<len(nums)-1;i++ {
        if nums[i]>nums[i+1] {
            down=up+1
        }
        if nums[i]<nums[i+1] {
            up=down+1
        }
    }
    if up > down {
        return up
    }
    return down
}
func wiggleMaxLengthXXX(nums []int) int {
    if len(nums) <= 0 {return 0}
    if len(nums) == 1 {return 1}
    num := 1
    for i:=1;i<len(nums)-1;i++ {
        if nums[i]>nums[i-1]&&nums[i]>nums[i+1] {
            num += 1
        } else if nums[i]<nums[i-1]&&nums[i]<nums[i+1] {
            num += 1
        }
    }
    return num+1
}
func main() {
    fmt.Println(wiggleMaxLength([]int{0,0}))
    fmt.Println(wiggleMaxLength([]int{1}))
    fmt.Println(wiggleMaxLength([]int{1,7,4,9,2,5}))
    fmt.Println(wiggleMaxLength([]int{1,17,5,10,13,15,10,5,16,8}))
    fmt.Println(wiggleMaxLength([]int{1,2,3,4,5,6,7,8,9}))
}
