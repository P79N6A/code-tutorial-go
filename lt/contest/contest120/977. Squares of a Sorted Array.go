package main

import "fmt"

/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.



Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

func main() {
    fmt.Println(sortedSquares([]int{-7,-3,2,3,11}))
    
}
func sortedSquares(A []int) []int {
    ans := make([]int,len(A))
    idx := len(A)-1
    l,r := 0,len(A)-1
    for idx >= 0 {
        ll := A[l]*A[l]
        rr := A[r]*A[r]
        if ll > rr {
            ans[idx]=ll
            l +=1
        } else {
            ans[idx]=rr
            r -=1
        }
        idx -=1
    }
    return ans
}
