package main

import (
        "sort"
        "fmt"
)
/*
Given an array of citations (each citation is a non-negative integer) of a researcher,
write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers
have at least h citations each, and the other N − h papers have no more than h citations each."

Example:

Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had
             received 3, 0, 6, 1, 5 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
Note: If there are several possible values for h, the maximum one is taken as the h-index.
*/

func hIndex1(citations []int) int {
        sort.Ints(citations)
        idx := 1
        for i:=len(citations)-1;i>=0;i-- {
                if citations[i] < idx {
                        break
                }
                idx += 1
        }
        return idx-1
}
func hIndex(citations []int) int {
        sort.Ints(citations)
        if len(citations) <= 0 {
                return 0
        }
        l,h := 0,len(citations)
        for l < h {
                m := (l+h)/2
                target := len(citations)-m
                //if citations[m] == target {
                //        return len(citations)-m
                //} else
                if citations[m] < target {
                        l = m+1
                } else {
                        h = m
                }
        }
        return len(citations)-h
}

func main() {
        //*
        fmt.Println(hIndex([]int{3,0,6,1,5}))
        fmt.Println(hIndex([]int{1,1,2,2}))
        fmt.Println(hIndex([]int{1,1}))
        fmt.Println(hIndex([]int{0,0}))
        fmt.Println(hIndex([]int{6,6,6,6,6,6}))
        //*/
        //fmt.Println(binsearch([]int{1,2,2,14,15},2))
        //fmt.Println(binsearch2([]int{1,2,2,14,15},2))
}
func binsearch2(nums []int,target int) int {
        l,h := 0,len(nums)
        for l<h {
                m := (l+h)/2
                fmt.Println(l,h,m)
                if nums[m] < target {
                        l=m+1
                } else {
                        h=m
                }
        }
        fmt.Println(l,h)
        return l
}
func binsearch(nums []int,target int) int {
        l,h := 0,len(nums)-1
        for l<=h {
                m := (l+h)/2
                fmt.Println(l,h,m)
                if nums[m] <= target {
                        l=m+1
                } else {
                        h=m-1
                }
        }
        fmt.Println(l,h)
        return h
}