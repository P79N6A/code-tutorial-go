package main

import "fmt"
/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


*/

func main() {
    fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
    fmt.Println(missingNumber([]int{0}))
    fmt.Println(missingNumber([]int{1}))
    fmt.Println(missingNumber([]int{1,2}))
    fmt.Println(missingNumber([]int{2,0}))
}
/*
问题:给定数组,1到n,缺失一个,问是哪个
思路:需要常数空间复杂度,由于也是1-n,就想到用数组本身做set. value对应的下标值乘以-1.
这样通过符号就能判断.
但corner case 还是很多的,比如* -1后0是不变的.
*/
func missingNumber(nums []int) int {
    // linear runtime complexity
    // constant extra space complexity
    extend := false
    haszero := false
    for i := 0; i < len(nums); i++ {
        idx := nums[i]
        if idx == 0 {haszero = true}
        if idx < 0 {idx *= -1}
        if idx >= len(nums) {
            extend = true
        } else {
            if nums[idx]>0 {
                nums[idx] *= -1
            }
        }
    }
    if !haszero {return 0}
    if !extend {return len(nums)}
    ret,reti := 0,0
    for i := 0; i < len(nums); i++ {
        if nums[i] >= ret { // 有可能0,x[x>0]都存在,找非0的.
            ret = nums[i]
            reti=i
        }
    }
    return reti
}