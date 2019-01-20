package main

import "fmt"

/*
A subarray A[i], A[i+1], ..., A[j] of A is said to be turbulent if and only if:

For i <= k < j, A[k] > A[k+1] when k is odd, and A[k] < A[k+1] when k is even;
OR, for i <= k < j, A[k] > A[k+1] when k is even, and A[k] < A[k+1] when k is odd.
That is, the subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

Return the length of a maximum size turbulent subarray of A.



Example 1:

Input: [9,4,2,10,7,8,8,1,9]
Output: 5
Explanation: (A[1] > A[2] < A[3] > A[4] < A[5])
Example 2:

Input: [4,8,12,16]
Output: 2
Example 3:

Input: [100]
Output: 1


Note:

1 <= A.length <= 40000
0 <= A[i] <= 10^9
*/

func main() {
    fmt.Println(maxTurbulenceSize([]int{9,4,2,10,7,8,8,1,9}))
    fmt.Println(maxTurbulenceSize([]int{4,8,12,16}))
    fmt.Println(maxTurbulenceSize([]int{100}))
    fmt.Println(maxTurbulenceSize([]int{0,1,1,0,1,0,1,1,0,0}))

}
func maxTurbulenceSize(A []int) int {
    return max(maxTurbulenceSize1(A),maxTurbulenceSize2(A))
}
func maxTurbulenceSize2(A []int) int {
    mans := 1
    ans := 1
    for i:=0;i<len(A)-1;i++ {
        if i%2 == 0 {
            if A[i] >A[i+1] {
                ans += 1
            } else {
                ans = 1
            }
        } else {
            if A[i] < A[i+1] {
                ans += 1
            } else {
                ans = 1
            }
        }
        mans = max(mans,ans)
    }
    return mans
}
func maxTurbulenceSize1(A []int) int {
    mans := 1
    ans := 1
    for i:=0;i<len(A)-1;i++ {
        if i%2 == 0 {
            if A[i] <A[i+1] {
                ans += 1
            } else {
                ans = 1
            }
        } else {
            if A[i] > A[i+1] {
                ans += 1
            } else {
                ans = 1
            }
        }
        mans = max(mans,ans)
    }
    return mans
}
func max(a,b int) int {
    if a>b {return a}
    return b
}
