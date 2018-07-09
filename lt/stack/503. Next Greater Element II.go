package main

import "fmt"

/*

Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

Example 1:
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2;
The number 2 can't find next greater number;
The second 1's next greater number needs to search circularly, which is also 2.
Note: The length of given array won't exceed 10000.
*/
func nextGreaterElements(nums []int) []int {
    stack := make([]int,2*len(nums))
    stackpos := make([]int,2*len(nums))
    idx := 0
    ret := make([]int,len(nums))
    for i:=0;i<len(ret);i++ {ret[i]=-1}
    for i:=0;i<2*len(nums);i++ {
        ci := i%len(nums)
        if idx ==0 || nums[ci]<=stack[idx-1] {
            stack[idx]=nums[ci]
            stackpos[idx]=ci
            idx += 1
        } else {
            for idx>0 {
                if stack[idx-1]<nums[ci] {
                    ret[stackpos[idx-1]]=nums[ci]
                    idx -= 1
                } else {
                    break
                }
            }
            stack[idx]=nums[ci]
            stackpos[idx]=ci
            idx += 1
        }
    }
    return ret
}
func main() {
    fmt.Println(nextGreaterElements([]int{1,2,3,4,3}))
    
}
