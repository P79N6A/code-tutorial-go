package main

import "fmt"
/*
Given an array of integers and an integer k,
you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

*/
func subarraySum(nums []int, k int) int {
        sumSet := make(map[int]int)
        sumSet[0]=1
        sum := 0
        total := 0
        for _,n := range nums {
                sum += n
                if _,ok := sumSet[sum-k];ok {
                        total += sumSet[sum-k]
                }
                sumSet[sum]+=1
        }
        return total
}
func main() {
        //fmt.Println(subarraySum([]int{1,1,1,1},2))
        //fmt.Println(subarraySum([]int{1},0))
        fmt.Println(subarraySum([]int{0,0,0,0},0))
}
