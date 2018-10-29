package main

import "fmt"

/*
Given a square array of integers A, we want the minimum sum of a falling path through A.

A falling path starts at any element in the first row, and chooses one element from each row.  The next row's choice must be in a column that is different from the previous row's column by at most one.



Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: 12
Explanation:
The possible falling paths are:
[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
The falling path with the smallest sum is [1,4,7], so the answer is 12.



Note:

1 <= A.length == A[0].length <= 100
-100 <= A[i][j] <= 100
*/

func main() {
  //  fmt.Println(minFallingPathSum([][]int{{1,2,3},{4,5,6},{7,8,9}}))
    fmt.Println(minFallingPathSum([][]int{{10,-98,44},{-20,65,34},{-100,-1,74}}))
}

func minFallingPathSum(A [][]int) int {
    if len(A)==0 {return 0}
    m,n := len(A),len(A[0])
    sum := make([]int,n)
    copy(sum,A[m-1])
    for i:=m-2;i>=0;i-- { // bottom-up
        sum2 := make([]int,n)
        for j:=0;j<len(sum);j++ {
            if j == 0 {
                sum2[j]=minf(sum[j],sum[j+1])
            } else if j == len(sum)-1 {
                sum2[j]=minf(sum[j],sum[j-1])
            } else {
                sum2[j]=minf(sum[j-1],minf(sum[j],sum[j+1]))
            }
            sum2[j] += A[i][j]
        }
        sum=sum2  // 空间优化，每次dp只跟上一次的相关
    }
    ans := sum[0]
    for i:=0;i<len(sum);i++ {
        if sum[i]<ans {ans = sum[i]}
    }
    return ans
}
func minFallingPathSum1(A [][]int) int {
    if len(A)==0 {return 0}
    min := 100001
    cache := make([][]int,0)
    for i:=0;i<len(A);i++ {
        cache = append(cache,make([]int,len(A[i])))
    }
    for i:=0;i<len(A[0]);i++ {
        x := solve(A,len(A),len(A[0]),0,i,0,&cache)
        if min > x {min = x}
    }
    return min
}

func solve(A [][]int,m,n int, row int, col int, sum int,cache *[][]int) int {
    if col < 0 || col>=n {return 100001}
    if row < 0 || row>=m {return 0}
    a := solve(A,m,n,row+1,col-1,sum+A[row][col],cache)
    b := solve(A,m,n,row+1,col,sum+A[row][col],cache)
    c := solve(A,m,n,row+1,col+1,sum+A[row][col],cache)
    (*cache)[row][col]=A[row][col]+minf(a,minf(b,c))
    return (*cache)[row][col]
}
func minf(a,b int) int {
    if a < b {return a}
    return b
}