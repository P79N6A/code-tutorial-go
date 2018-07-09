package main

import (
    "fmt"
)

/*

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

*/
func longestConsecutive(nums []int) int {
    root := make(map[int]int)
    for i := 0;i < len(nums);i++ {
        root[nums[i]]=nums[i]
        if _,ok := root[nums[i]+1];ok {
            root[nums[i]]=nums[i]+1
        }
        if _,ok := root[nums[i]-1];ok {
            root[nums[i]-1]=nums[i]
        }
    }
    max := 0
    cache := make(map[int]int)
    for k,_ := range root {
        dfs(k,root,cache,&max)
    }
    return max
}
// 联通图的大小
func dfs(i int,r map[int]int,cache map[int]int, max *int) int {
    if cache[i] > 0 {
        return cache[i]
    }
    // 输出位置点的check
    if r[i] == i {
        cache[i]=1
        if *max < 1 {
            *max = 1
        }
        return 1
    }
    l := dfs(r[i],r,cache,max)
    if *max < l+1 {
        *max = l+1
    }
    cache[i]=l+1
    return l+1
}
func main() {
    fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
