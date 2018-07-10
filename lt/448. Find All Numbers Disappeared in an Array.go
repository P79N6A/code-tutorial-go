package main

import "fmt"
/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

*/
func findDisappearedNumbers(nums []int) []int {
        for i:=0;i<len(nums);i++ {
                idx := nums[i]
                if idx < 0 {idx *= -1}
                if nums[idx-1] > 0 {
                        nums[idx-1] *= -1
                }
        }
        ret := make([]int,0)
        for i:=0;i<len(nums);i++ {
                if nums[i] > 0 {
                        ret = append(ret,i+1)
                }
        }
        return ret
}

func main() {
        fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}
