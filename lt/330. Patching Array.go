package main

import "fmt"

/*
Given a sorted positive integer array nums and an integer n, add/patch elements to the array such that any number in range [1, n] inclusive can be formed by the sum of some elements in the array. Return the minimum number of patches required.

Example 1:

Input: nums = [1,3], n = 6
Output: 1
Explanation:
Combinations of nums are [1], [3], [1,3], which form possible sums of: 1, 3, 4.
Now if we add/patch 2 to nums, the combinations are: [1], [2], [3], [1,3], [2,3], [1,2,3].
Possible sums are 1, 2, 3, 4, 5, 6, which now covers the range [1, 6].
So we only need 1 patch.
Example 2:

Input: nums = [1,5,10], n = 20
Output: 2
Explanation: The two patches can be [2, 4].
Example 3:

Input: nums = [1,2,2], n = 5
Output: 0
*/

func main() {
    fmt.Println(minPatches([]int{1,5,10},20))
    
}
func minPatches(nums []int, n int) int {
    add := 0
    miss := 1 // 缺失的数字
    i := 0
    for miss <= n {
        if i < len(nums) && nums[i] <= miss {
            // 如果有一个小于确实值的，则往前移动就行了，表达的数字叠加nums[i]
            // 比如miss可以表达1-7， 来了一个5，则范围肯定变成了1-(5+7)了
            miss += nums[i]
            i += 1
        } else {
            /*
            为什么加上miss？
            想想如果给的nums是空的，该如何作答？
            都是2的幂，
            */
            miss += miss
            add += 1
        }
    }
    return add
}