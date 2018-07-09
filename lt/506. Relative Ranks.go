package main

import (
        "sort"
        "strconv"
        "fmt"
)
/*
Given scores of N athletes, find their relative ranks and the people with the top three highest scores,
who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
For the left two athletes, you just need to output their relative ranks according to their scores.
Note:
N is a positive integer and won't exceed 10,000.
All the scores of athletes are guaranteed to be unique.
*/

func findRelativeRanks(nums []int) []string {
        rnums := make([]int,len(nums))
        copy(rnums,nums)
        sort.Ints(rnums)
        ret := make([]string,0)
        for _,n := range nums {
                idx := binarysearch(rnums,n)
                if idx == len(nums)-1 {
                        ret = append(ret,"Gold Medal")
                } else if idx == len(nums)-2 {
                        ret = append(ret,"Silver Medal")
                } else if idx == len(nums)-3 {
                        ret = append(ret,"Bronze Medal")
                } else {
                        ret = append(ret,strconv.Itoa(len(nums)-idx))
                }
        }
        return ret
}
func binarysearch(nums []int,t int) int {
        l,h := 0,len(nums)
        for l < h {
                m := (l+h)/2
                if nums[m] == t {
                        return m
                } else if nums[m] < t {
                        l = m + 1
                } else {
                        h = m
                }
        }
        return h
}

func main() {
        fmt.Println(findRelativeRanks([]int{5,4,3,2,1}))



}
