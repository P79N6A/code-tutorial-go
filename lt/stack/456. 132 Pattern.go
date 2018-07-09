package main

import (
    "math"
    "fmt"
)

/*
Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence ai, aj, ak such that i < j < k and ai < ak < aj. Design an algorithm that takes a list of n numbers as input and checks whether there is a 132 pattern in the list.

Note: n will be less than 15,000.

Example 1:
Input: [1, 2, 3, 4]

Output: False

Explanation: There is no 132 pattern in the sequence.
Example 2:
Input: [3, 1, 4, 2]

Output: True

Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
Example 3:
Input: [-1, 3, 2, 0]

Output: True

Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].

还是不太好想，一个单调栈能干什么？找到第一个，或者是否有大于/小于 某个数的元素。

min是为了找 1，3的， 然后通过单调栈看是否有2存在


*/
func find132pattern(nums []int) bool {
    if len(nums) < 3 {return false}
    min := make([]int,len(nums))
    mx := math.MaxInt64
    for i:=0;i<len(nums);i++ {
        if nums[i] < mx {
            mx = nums[i]
        }
        min[i]=mx
    }
    stack := make([]int,0)
    // 从后往前
    for i:=len(nums)-1;i>=0;i-- {
        if nums[i]>min[i] { // 1-3 关系满足
            for len(stack)> 0 && stack[len(stack)-1] <= min[i] { // 1-2关系满足
                // 如果栈顶元素比当前最小值，也就是1 都小了，那就是没必要留下来了
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 && stack[len(stack)-1] < nums[i]  { // 是否是3-2关系
                // 如果发现了一个3-2关系
                return true
            }
            // 当前关系已经是1-3-4了，3入栈 递减栈关系，入栈，看下一个是否符合逻辑
            stack = append(stack,nums[i])
        }
    }
    return false
}

func main() {
    fmt.Println(find132pattern([]int{6,11,12,3,4,5}))
}