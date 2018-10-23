package main

import (
    "strings"
    "fmt"
)

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/

func main() {
    fmt.Println(solveNQueens(8))
    
}
func solveNQueens(n int) [][]string {
    ans := make([][]string,0)
    l := strings.Repeat(".",n)
    bs := make([][]byte,0)
    for i:=0;i<n;i++ {
        bs = append(bs,[]byte(l))
    }
    ij := make(map[int]bool)  // 左下=>右上方向visit记录
    i_j := make(map[int]bool) // 左上=>右下方向visit记录
    col := make(map[int]bool) //第几列visit记录
    solveRow(0,n,&bs,&ans,col,ij,i_j)
    return ans
}
func solveRow(row int,n int,bs *[][]byte,ans *[][]string,col,ij,i_j map[int]bool) {
    if row==n {
        a := make([]string,0)
        for _,b := range *bs {
            a = append(a,string(b))
        }
        *ans = append(*ans,a)
        return
    }
    for i:=0;i<n;i++ {
        if col[i] {continue}
        if ij[row+i] {continue}
        if i_j[row-i]{continue}
        ij[row+i]=true
        col[i]=true
        i_j[row-i]=true
        (*bs)[row][i]='Q'
        solveRow(row+1,n,bs,ans,col,ij,i_j)
        (*bs)[row][i]='.'
        ij[row+i]=false
        col[i]=false
        i_j[row-i]=false
    }
    return
}