package main

import "fmt"
/*
Let's call any (contiguous) subarray B (of A) a mountain if the following properties hold:

B.length >= 3
There exists some 0 < i < B.length - 1 such that B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
(Note that B could be any subarray of A, including the entire array A.)

Given an array A of integers, return the length of the longest mountain.

Return 0 if there is no mountain.



Example 1:

Input: [2,1,4,7,3,2,5]
Output: 5
Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
Example 2:

Input: [2,2,2]
Output: 0
Explanation: There is no mountain.


Note:

0 <= A.length <= 10000
0 <= A[i] <= 10000
*/
func longestMountain(A []int) int {
        mountain := make([]int,0)
        down := make([]int,0)
        if len(A) < 2 {
                return 0
        }
        if A[0] > A[1] {
                mountain = append(mountain,0)
        } else if A[0]<A[1] {
                down = append(down,0)
        }
        for i:=1;i<len(A)-1;i++ {
                if A[i]>A[i-1] && A[i] > A[i+1] {
                        mountain = append(mountain,i)
                }
                if A[i]<=A[i-1] && A[i] <= A[i+1] {
                        down = append(down,i)
                }
        }
        if len(A) > 3 && A[len(A)-1] > A[len(A)-2] {
                mountain = append(mountain,len(A)-1)
        } else if len(A) >= 3 && A[len(A)-1] < A[len(A)-2]  {
                down = append(down,len(A)-1)
        }
        fmt.Println(mountain)
        fmt.Println(down)
        ret := 0
        for _,m := range mountain {
                idx := binarysearch(down,m)
                c := -1
                if idx >= len(down) {
                //        c = m - down[len(down)-1]
                } else if idx <= 0{
                //        c = down[0]-m
                } else {
                        c = down[idx] - down[idx-1]
                }
                if c+1 > ret {
                        ret = c+1
                }
        }
        return ret
}
func binarysearch(nums []int,t int) int {
        l,h := 0,len(nums)
        for l < h {
                m := (l+h)/2
                if nums[m] == t {
                        return t
                } else  if nums[m] < t {
                        l = m + 1
                } else {
                        h = m
                }
        }
        return h
}

func main() {
        //fmt.Println(longestMountain([]int{2,1,4,7,3,2,5}))
        //fmt.Println(longestMountain([]int{2,2,2}))
        //fmt.Println(longestMountain([]int{3,2}))
        //fmt.Println(longestMountain([]int{1,1,0,0,1,0}))
        //fmt.Println(longestMountain([]int{0,1,0}))
        fmt.Println(longestMountain([]int{0,0,1,0,0,1,1,1,1,1}))
        //fmt.Println(longestMountain([]int{0,1,2,3,4,5,4,3,2,1,0}))
}
