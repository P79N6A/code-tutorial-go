package main

import "sort"

/*
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.



Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]
Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9
*/

/*
分析问题，binarySearch+greedy，找第一个比B[i]大的，这样能产生advantage，如果没有，就用最小的顶
*/
func advantageCount(A []int, B []int) []int {
    ret := make([]int,0)
    alen := len(A)
    sort.Ints(A)
    for i:=0;i<len(B);i++ {
        ok,idx := upper(A,alen,B[i])
        if !ok {
            idx = 0
        }
        //fmt.Println(A,alen,B[i],idx,ok)
        ret = append(ret,A[idx])
        copy(A[idx:],A[(idx+1):])
        alen -=1
    }
    return ret
}
func upper(A []int,alen int, target int) (bool,int) {
    l,r := 0,alen
    for l<r {
        m := (l+r)/2
        if A[m]<=target {
            l = m+1
        } else {
            r=m
        }
    }
    if r >= (alen) {return false,r}
    return true,r
}
func main() {
    
}
