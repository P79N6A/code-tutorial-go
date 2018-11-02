package main

import "fmt"

/*
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:

Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
Example 2:

Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.
Note:

You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
Try to solve it in linear time/space.

*/

func main() {
    fmt.Println(maximumGap([]int{3,6,9,1}))
    
}
func maximumGap(nums []int) int {
    nums = radixSort(nums)
    maxgap := 0
    for i:=1;i<len(nums);i++ {
        maxgap = max(maxgap,nums[i]-nums[i-1])
    }
    return maxgap
}
func radixSort(nums []int) []int {
    // reference : https://www.cs.usfca.edu/~galles/visualization/RadixSort.html
    if len(nums) < 2 {return nums}
    m := nums[0]
    for _,n := range nums {m = max(m,n)}
    radix := 10
    exp := 1
    aux := make([]int,len(nums))
    for m/exp >0 { // 从小到大遍历10进制位
        cnt := make([]int,radix) // 记录当前遍历的位，余数出现次数
        for i:=0;i<len(nums);i++ {
            cnt[nums[i]/exp%10]+=1 // 指定exp位数的10余数
        }
        for i:=1;i<len(cnt);i++ {
            // 叠加上，这样radix的各个位置的最后一位的下标就是有意义的
            // 想下，在收集的过程中，8肯定在9的前边，那9的位置肯定都是8的后边
            cnt[i]+=cnt[i-1]
        }
        for i:=len(nums)-1;i>=0;i-- {
            // 从后往前收集;由于收集数组aux是每次都会覆盖的，所以不用每次新建
            cnt[nums[i]/exp%10] -= 1
            aux[cnt[nums[i]/exp%10]] = nums[i] // 从后往前，按照radix 填充到指定位置
        }
        for i:=0;i<len(aux);i++ {
            nums[i]=aux[i] //  放回去
        }
        exp *= 10
    }
    return nums
}
func max(a,b int) int {
    if a < b {return b}
    return a
}
