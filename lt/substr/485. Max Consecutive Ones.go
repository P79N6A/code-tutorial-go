package main

import "fmt"
/*

Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000

*/
func findMaxConsecutiveOnes(nums []int) int {
        left,right,zero := 0,0,0
        res := 0
        for ;right < len(nums);right++ {
                if nums[right] == 0 {
                        zero += 1
                }
                for zero > 0 {
                        if nums[left] == 0 {
                                zero -= 1
                        }
                        left += 1
                }
                if right - left +1 > res {
                        res = right - left + 1
                }
        }
        return res
}

func main() {
        fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,0}))
}
