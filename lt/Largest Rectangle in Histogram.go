package main

import "fmt"

/*
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].



The largest rectangle is shown in the shaded area, which has area = 10 unit.



Example:

Input: [2,1,5,6,2,3]
Output: 10
*/

func largestRectangleArea(heights []int) int {
    stack := make([]int,len(heights))
    idx := 0
    max := 0
    // 单调栈，看到的永远是左侧第一个比自己小的
    // i 确定了右边界,stack中比自己小的那个确定了左边界!!!!
    if len(heights) > 0 {
        heights = append(heights,0)
    }
    for i:=0;i<len(heights);i++ {
        if idx == 0 || heights[i] > heights[stack[idx-1]] {
            stack[idx]=i
            idx += 1
        } else {
            cur := stack[idx-1]
            idx -= 1
            if idx == 0 {
                if max < heights[cur]*(i) {
                    max = heights[cur] * (i)
                }
            } else {
                if max < heights[cur]*(i-stack[idx-1]-1) {
                    max = heights[cur] * (i - stack[idx-1]-1)
                }
            }
            i-=1
        }
    }
    return max
}

func main() {
    fmt.Println(largestRectangleArea([]int{}))
    fmt.Println(largestRectangleArea([]int{2,1,2}))
    fmt.Println(largestRectangleArea([]int{4,2}))
    fmt.Println(largestRectangleArea([]int{1,3,2,1,5,6,4}))
    fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))

}
