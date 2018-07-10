package main

import (
        "sort"
        "fmt"
)
/*
Given an integer array, return the k-th smallest distance among all the pairs.
The distance of a pair (A, B) is defined as the absolute difference between A and B.

Example 1:
Input:
nums = [1,3,1]
k = 1
Output: 0
Explanation:
Here are all the pairs:
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
Then the 1st smallest distance pair is (1,1), and its distance is 0.
Note:
2 <= len(nums) <= 10000.
0 <= nums[i] < 1000000.
1 <= k <= len(nums) * (len(nums) - 1) / 2.
*/
func smallestDistancePair(nums []int, k int) int {
        sort.Ints(nums)
        l,r := 0,nums[len(nums)-1]
        for l<r {
                m := (l+r)/2
                counter := 0
                for i:=0;i<len(nums);i++ {
                        for j:=i+1;j<len(nums);j++ {
                                if nums[j]-nums[i]<m {
                                        counter += 1
                                }
                        }
                }
                if counter < k {
                        l = m+1
                } else {
                        r = m
                }
        }
        return r-1
}
func main() {
        fmt.Println(smallestDistancePair([]int{1,3,1},1))
}
