package main

import (
    "fmt"
    "math"
)

/*
We have two integer sequences A and B of the same non-zero length.

We are allowed to swap elements A[i] and B[i].  Note that both elements are in the same index position in their respective sequences.

At the end of some number of swaps, A and B are both strictly increasing.  (A sequence is strictly increasing if and only if A[0] < A[1] < A[2] < ... < A[A.length - 1].)

Given A and B, return the minimum number of swaps to make both sequences strictly increasing.  It is guaranteed that the given input always makes it possible.

Example:
Input: A = [1,3,5,4], B = [1,2,3,7]
Output: 1
Explanation:
Swap A[3] and B[3].  Then the sequences are:
A = [1, 3, 5, 7] and B = [1, 2, 3, 4]
which are both strictly increasing.
Note:

A, B are arrays with the same length, and that length will be in the range [1, 1000].
A[i], B[i] are integer values in the range [0, 2000].
*/
func main() {
    fmt.Println(minSwap([]int{1,3,5,4},[]int{1,2,3,7}))
    //[0,3,5,8,9]
    //[2,1,4,6,9]
    fmt.Println(minSwap([]int{0,3,5,8,9},[]int{2,1,4,6,9}))
    //[4,10,13,14,17,19,21,24,26,27,28,29,34,37,38,42,44,46,48,51,52,53,54,57,58,59,64,65,66,67,71,73,75,76,80,81,82,83,86,88,89,90,95,97,98,99,101,105,106,108,109,110,111,112,115,119,121,122,123,124,125,126,127,128,129,130,131,133,136,138,143,145,147,149,150,153,158,160,163,164,165,167,168,169,172,173,174,176,178,179,183,184,186,188,189,192,193,194,198,200]
    //[0,1,3,5,6,7,11,13,15,16,17,21,37,39,41,42,43,45,47,50,53,55,56,57,64,66,67,68,69,70,71,72,74,75,76,77,79,80,87,88,89,95,96,97,98,100,101,105,106,107,108,112,113,115,116,118,119,122,124,125,126,127,128,131,135,136,137,138,139,140,144,145,148,150,151,154,159,160,161,162,163,167,168,170,171,174,176,178,179,180,181,185,187,189,190,191,192,198,199,200]
    fmt.Println(minSwap2([]int{
        4,10,13,14,17,19,21,24,26,27,28,29,34,37,38,42,44,46,48,51,52,53,54,57,58,59,64,65,66,67,71,73,75,76,80,81,82,83,86,88,89,90,95,97,98,99,101,105,106,108,109,110,111,112,115,119,121,122,123,124,125,126,127,128,129,130,131,133,136,138,143,145,147,149,150,153,158,160,163,164,165,167,168,169,172,173,174,176,178,179,183,184,186,188,189,192,193,194,198,200,
    },[]int{
        0,1,3,5,6,7,11,13,15,16,17,21,37,39,41,42,43,45,47,50,53,55,56,57,64,66,67,68,69,70,71,72,74,75,76,77,79,80,87,88,89,95,96,97,98,100,101,105,106,107,108,112,113,115,116,118,119,122,124,125,126,127,128,131,135,136,137,138,139,140,144,145,148,150,151,154,159,160,161,162,163,167,168,170,171,174,176,178,179,180,181,185,187,189,190,191,192,198,199,200,
    }))

}
func minSwap2(A []int, B []int) int {
    n := len(A)
    keep := make([]int,n)
    swap := make([]int,n)
    for i:=0;i<n;i++ {keep[i]=math.MaxInt64}
    for i:=0;i<n;i++ {swap[i]=math.MaxInt64}
    keep[0]=0
    swap[0]=1
    for i:=1;i<len(A);i++ {
        if A[i-1]<A[i]&&B[i-1]<B[i] {
            // 保持状态，不交换
            keep[i]=keep[i-1]
            // A[i] B[i]想交换，必须得交换A[i-1] B[i-1]
            swap[i]=swap[i-1]+1
        }
        if A[i-1]<B[i]&&B[i-1]<A[i] {
            // 由于两种情况可能同时存在，所以第二次需要使用min
            // 如果A[i] B[i]不交换，这种情况可以交换A[i-1] B[i-1]
            keep[i]=min(keep[i],swap[i-1])
            // 如果A[i] B[i]交换，这种情况那就不要交换A[i-1] B[i-1]
            swap[i]=min(swap[i-1],keep[i-1]+1)
        }
    }
    return min(keep[n-1],swap[n-1])
}
func min(a,b int) int {
    if a > b {return b}
    return a
}
func minSwap(A []int, B []int) int {
    /*
    keep指的是不交换第i个元素的时候的最小交换次数
    swap指的是交换第i个元素的时候的最小交换次数
    */
    keep,swap := 0,1
    for i:=1;i<len(A);i++ {
        if A[i-1]>=B[i] || B[i-1]>=A[i] {
            // A[i] B[i] 不能交换的情况
            // keep继续保持不交换的zhuangtai
            keep = keep
            // 这种情况如果想交换A[i] B[i] 则必须也交换 A[i-1] B[i-1]
            swap = swap+1
        } else if A[i-1] >= A[i] || B[i-1] >= B[i] {
            // A[i] B[i] 必须交换的情况
            k := keep
            // keep 等于A[i-1] B[i-1]交换，因为i-1的两个交换，那i不用交换也就合理了
            keep = swap
            // swap 等于A[i-1] B[i-1]不交换， A[i]B[i] 交换，
            swap = k+1
        } else {
            // 可以交换可以不交换，那就取较小的即可。
            if keep > swap {keep = swap}
            swap = keep + 1
        }
    }
    if swap > keep {return keep}
    return swap
}
func minSwap1(A []int, B []int) int {
    swap := len(A)+1
    solve(A,B,1,len(A),0,&swap)
    if swap == len(A)+1 {return -1}
    return swap
}
func solve(A,B []int,idx,len int,num int, swap *int) {
    if idx == len {
        fmt.Println(idx,num)
        if increase(A) && increase(B) {if num < *swap {*swap = num}}
        return
    }
    if A[idx] > A[idx-1] && B[idx] > B[idx-1] {
        solve(A,B,idx+1,len,num,swap)
    }
    if B[idx]>A[idx-1] && A[idx]>B[idx-1]{
        A[idx],B[idx]=B[idx],A[idx]
        solve(A,B,idx+1,len,num+1,swap)
        A[idx],B[idx]=B[idx],A[idx]
    }
}
func increase(A []int) bool {
    for i:=1;i<len(A);i++ {if A[i]<=A[i-1] {return false}}
    return true
}

