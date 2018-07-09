package main

import "fmt"

/*

Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

*/
// a.找连续子数组的和等于指定的数字[可以有负数]
func subarraySum(nums []int, k int) int {
    //key 曾经出现过的sum；value 是个数
    sumset := make(map[int]int)
    sum,cnt := 0,0
    sumset[0]=1
    for i:=0;i<len(nums);i++ {
        sum += nums[i] // 新的sum值
        cnt += sumset[sum-k] //不算当前的，看历史是否存在sum-k
        sumset[sum]+=1
    }
    return cnt
}
func main() {
    fmt.Println(subarraySum([]int{1,1,1},2))
}
