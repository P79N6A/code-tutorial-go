package main

import "fmt"

/*
Given a sorted array, two integers k and x, find the k closest elements to x in the array.
The result should also be sorted in ascending order.
If there is a tie, the smaller elements are always preferred.

Example 1:
Input: [1,2,3,4,5], k=4, x=3
Output: [1,2,3,4]
Example 2:
Input: [1,2,3,4,5], k=4, x=-1
Output: [1,2,3,4]
Note:
The value k is positive and will always be smaller than the length of the sorted array.
Length of the given array is positive and will not exceed 104
Absolute value of elements in the array and x will not exceed 104
UPDATE (2017/9/19):
The arr parameter had been changed to an array of integers (instead of a list of integers).
Please reload the code definition to get the latest changes.
*/
func findClosestElements(arr []int, k int, x int) []int {
    idx := lowerBound(arr,x)
    start := idx - k
    if start < 0 {start = 0}
    end := idx + k
    if end >= len(arr) {
        end = len(arr)-1
    }
    for start <= end {
        if end - start <= k-1 {
            return arr[start:end+1]
        }
        if x - arr[start] > arr[end]-x {
            start += 1
        } else {
            end -= 1
        }
    }
    return nil
}

func lowerBound(arr []int,x int) int {
    l,r := 0,len(arr)
    for l < r {
        m := (l+r)/2
        if arr[m] < x {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}

func main() {
    //fmt.Println(findClosestElements([]int{1,2,3,4,5},3,3))
    fmt.Println(findClosestElements([]int{1},1,1))
    //fmt.Println(findClosestElements([]int{1,2,3,4,5},4,-1))

}
