package main

import (
    "fmt"
)

func main() {
  //  fmt.Println(minFallingPathSum([][]int{{1,2,3},{4,5,6},{7,8,9}}))
    fmt.Println(minFallingPathSum([][]int{{10,-98,44},{-20,65,34},{-100,-1,74}}))
}

func minFallingPathSum(A [][]int) int {
    if len(A)==0 {return 0}
    m,n := len(A),len(A[0])
    sum := make([]int,n)
    copy(sum,A[m-1])
    for i:=m-2;i>=0;i-- {
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
        sum=sum2
    }
    ans := sum[0]
    for i:=0;i<len(sum);i++ {
        if sum[i]<ans {
            ans = sum[i]
        }
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