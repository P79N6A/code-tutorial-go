package main

import "fmt"
/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

*/

func longestConsecutive(nums []int) int {
        if len(nums) <= 0 {return 0}
        nseq := make(map[int]int)
        for i:=0;i<len(nums);i++ {
                if _,ok := nseq[nums[i]];!ok {
                        nseq[nums[i]] = nums[i]
                }
                if _,ok := nseq[nums[i]-1];ok {
                        nseq[nums[i]-1] = nums[i]
                }
                if _,ok := nseq[nums[i]+1];ok {
                        nseq[nums[i]] = nums[i]+1
                }
        }
        cache := make(map[int]int)
        max := 1
        for k,_ := range nseq {
                dfs(nseq,k,cache,&max)
        }
        return max
}
func dfs(nseq map[int]int,i int,cache map[int]int, max *int) int {
        if cache[i] != 0 {
                return cache[i]
        }
        if _,ok := nseq[i];!ok {
                return 0
        }
        if nseq[i] == i {
                cache[i]=1
                if 1 > *max {*max=1}
                return 1
        }
        n := 1 + dfs(nseq,i+1,cache,max)
        cache[i]=n
        if n > *max {*max=n}
        return n
}

func main() {
        //fmt.Println(longestConsecutive([]int{100,4,200,1,2,3}))
        fmt.Println(longestConsecutive([]int{0,-1}))
}
