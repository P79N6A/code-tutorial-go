package main

import "fmt"

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

func main() {
    fmt.Println(findMedianSortedArrays([]int{1,3},[]int{2}))
    fmt.Println(findMedianSortedArrays([]int{1,2},[]int{3,4}))
    fmt.Println(findKthInTwoSortArray([]int{1,3},[]int{2},2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m,n := len(nums1),len(nums2)
    return float64(findKthInTwoSortArray(nums1,nums2,(m+n+1)/2)+findKthInTwoSortArray(nums1,nums2,(m+n+2)/2))/2
}
func findKthInTwoSortArray(A,B []int,k int) int {
    // k是右开区间边界，比下标大了一个
    if len(A) <= 0 {return B[k-1]}
    if len(B) <= 0 {return A[k-1]}
    if k==1 {return min(A[0],B[0])}
    // 右边界开区间，所以是len(A)，下边i-1保证肯定有效。因为上边已经判断len(A)>0
    // k/2 有可能 > len(A)
    i := min(len(A),k/2)
    j := min(len(B),k/2)
    if A[i-1]>B[j-1] {
        return findKthInTwoSortArray(A,B[j:],k-j)
    }
    return findKthInTwoSortArray(A[i:],B,k-i)
}
func min(a,b int)int {
    if a<b {return a}
    return b
}