package main

import "fmt"

/*
You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2. Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.

Example 1:
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
Example 2:
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
Note:
All elements in nums1 and nums2 are unique.
The length of both nums1 and nums2 would not exceed 1000.
*/

func main() {
    fmt.Println(nextGreaterElement([]int{4,1,2},[]int{1,3,4,2}))
    fmt.Println(nextGreaterElement([]int{2,4},[]int{1,2,3,4}))
}
func nextGreaterElement(findNums []int, nums []int) []int {
    stack := make([]int,0)
    next := make(map[int]int)
    for i:=0;i<len(nums);i++ {
        /*
        单调递减栈；入栈之前可以得到【栈顶右侧第一个比栈顶大的元素】
        注意先处理出栈，再处理入栈；代码更简洁
        */
        for len(stack)>0 && nums[i]>stack[len(stack)-1] {
            next[stack[len(stack)-1]]=nums[i]
            stack = stack[:(len(stack)-1)]
        }
        stack = append(stack,nums[i])
    }
    ans := make([]int,len(findNums))
    for i:=0;i<len(findNums);i++ {
        if _,ok := next[findNums[i]];ok {
            ans[i]=next[findNums[i]]
        } else {
            ans[i]=-1
        }
    }
    return ans
}