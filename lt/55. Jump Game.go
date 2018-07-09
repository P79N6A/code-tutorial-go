package main

import "fmt"
/*
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.

*/

func canJump(nums []int) bool {
        /*
        l就是起始点
        r就是最远能达到的距离
        */
        l,r,nlen := 0,0,len(nums)
        if nlen <= 1 {return true}
        for {
                if l + nums[l] > r {
                        r = l + nums[l]
                }
                l += 1
                if r >= nlen-1 {
                        // 如果r到达最后一位了,canjump
                        return true
                }
                if l > r {
                        //如果l已经到了r,说明无法向前了
                        break
                }
        }
        return false
}

func main() {
        fmt.Println(canJump([]int{2,0,0}))
        fmt.Println(canJump([]int{2,3,1,1,4}))
        fmt.Println(canJump([]int{3,2,1,0,4}))
}
