package main

import "fmt"

/*
Given an array which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays. Write an algorithm to minimize the largest sum among these m subarrays.

Note:
If n is the length of array, assume the following constraints are satisfied:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
Examples:

Input:
nums = [7,2,5,10,8]
m = 2

Output:
18

Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
*/
func splitArray(nums []int, m int) int {
    left,right := 0,0
    //计算左右边界，分成1组，或者分成len(nums)组形成最大，最小值
    for i:=0;i<len(nums);i++ {
        if left < nums[i] {left = nums[i]}
        right += nums[i]
    }
    for left < right {
        mid := (left + right)/2
        // 如果还能拆分成功，说明mid太大了，right左移缩小mid
        if canSplit(nums,m,mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
func canSplit(nums []int,m int, maxsum int) bool {
    //是否能切分？ nums分成m组，每组最大值不能超过maxsum
    sum,group := 0,1 // group初始值1，因为默认已有一组
    for i:=0;i<len(nums);i++ {
        sum += nums[i]
        if sum > maxsum {
            sum = nums[i]
            group += 1
            if group > m {return false}
        }
    }
    return true
}

func main() {
    fmt.Println(splitArray([]int{7,2,5,10,8},2))
    
}
