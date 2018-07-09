package main

import (
    "sort"
    "fmt"
)

/*
Given an integer array nums, return the number of range sums that lie in [lower, upper] inclusive.
Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j (i ≤ j), inclusive.

Note:
A naive algorithm of O(n2) is trivial. You MUST do better than that.

Example:

Input: nums = [-2,5,-1], lower = -2, upper = 2,
Output: 3
Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective sums are: -2, -1, 2.

*/

func countRangeSum(nums []int, lower int, upper int) int {
    sum := make([]int,len(nums)+1)
    for i:=0;i<len(nums);i++ {sum[i+1] = sum[i] + nums[i]}
    return mergeSort(sum,0,len(sum),lower,upper)

}
func mergeSort(sum []int, start,end int,lower,upper int) int {
    if end - start <= 1 {
        return 0
    }
    mid := (start + end)/2
    count := mergeSort(sum,start,mid,lower,upper) + mergeSort(sum,mid,end,lower,upper)
    m,n := mid,mid
    for i:=start;i<mid;i++ {
        for m < end && sum[m]-sum[i]<lower {
            m+=1
        }
        for n<end && sum[n]-sum[i]<=upper  {
            n+=1
        }
        count += n-m
    }
    // 应该写mergesort 子程序
    sort.Ints(sum[start:end])
    return count
}

func main() {
    fmt.Println(countRangeSum([]int{-2,5,-1},-2,2))
    
}
