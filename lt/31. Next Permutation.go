package main

import (
    "fmt"
    "sort"
)

/*

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) []int {
    // find the first down
    idx := len(nums)-1
    for ;idx >= 1&&nums[idx] <= nums[idx-1];idx-- {}
    idx -= 1
    if idx < 0 {
        sort.Ints(nums)
    } else {
        i := len(nums)-1
        for ;i>=0&&nums[i]<=nums[idx];i-- {}
        if i >=0 {
            nums[i],nums[idx]=nums[idx],nums[i]
            sort.Ints(nums[idx+1:])
        }
    }
    return nums
}

func main() {
    fmt.Println(nextPermutation([]int{1,2,3}))
    fmt.Println(nextPermutation([]int{3,2,1}))
    fmt.Println(nextPermutation([]int{1,1,5}))
    fmt.Println(nextPermutation([]int{3,5,4,2,1}))
    fmt.Println(nextPermutation([]int{1,3,5,4,2,1}))
    fmt.Println(nextPermutation([]int{1,5,1}))// 5,1,1
    fmt.Println(nextPermutation([]int{1,1}))// 5,1,1
    fmt.Println(nextPermutation([]int{5,1,1}))// 5,1,1
    fmt.Println(nextPermutation([]int{3,4,5,3,2}))// 5,1,1

}
