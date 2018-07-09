package main

import "fmt"
/*
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1
Note:

3 <= A.length <= 10000
0 <= A[i] <= 10^6
A is a mountain, as defined above.
*/
func peakIndexInMountainArray(A []int) int {
        if len(A) <= 1{return 0}
        l,r := 0,len(A)
        for l < r {
                m := (l+r)/2
                if m == 0 {
                       if  A[m]>A[m+1] {
                               return m
                       } else {
                               l = m+1
                       }
                } else  if m == len(A)-1{
                        if A[m] > A[m-1] {
                                return m
                        } else {
                                r = m
                        }
                } else {
                        if A[m] > A[m+1] && A[m] > A[m-1] {
                                return m
                        } else if A[m]<A[m+1]&&A[m]>A[m-1] {
                                l = m+1
                        } else {
                                r = m
                        }
                }
        }
        return r
}


func main() {
        fmt.Println(peakIndexInMountainArray([]int{0,2,1,0}))
        fmt.Println(peakIndexInMountainArray([]int{0,1,0}))
        fmt.Println(peakIndexInMountainArray([]int{0,0}))
        fmt.Println(peakIndexInMountainArray([]int{0}))
}
