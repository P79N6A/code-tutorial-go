package main

import "fmt"

/*
Your are given an array of positive integers nums.

Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.

Example 1:
Input: nums = [10, 5, 2, 6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
Note:

0 < nums.length <= 50000.
0 < nums[i] < 1000.
0 <= k < 10^6.
*/

func main() {
    fmt.Println(numSubarrayProductLessThanK([]int{10,5,2,6},100))
    fmt.Println(numSubarrayProductLessThanK([]int{10,5,100,2,6},100))
    fmt.Println(numSubarrayProductLessThanK([]int{1,2,3},0))
}
func numSubarrayProductLessThanK(nums []int, k int) int {
    l,r := 0,1
    ans := 0
    product := nums[0]
    for ;l<len(nums);l++ {
        for ;r<len(nums)&&product<k;r++ {
            product *=nums[r]
        }
        if l < r {
            if product >= k {
                ans += r - l - 1
            } else {
                ans += r - l
            }
        }
        product /= nums[l]
    }
    return ans
}
