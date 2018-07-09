package graph

import "fmt"

/*

Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.

*/
func longestIncreasingPath(matrix [][]int) int {
    if len(matrix)<=0 {return 0}
    m := len(matrix)
    n := len(matrix[0])
    max := 0
    cache := make([][]int,0)
    for i:=0;i<m;i++ {
        cache = append(cache,make([]int,n))
    }
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            x := dfs(matrix,i,j,m,n,cache)
            if x > max {
                max = x
            }
        }
    }
    return max
}
var dir [][]int = [][]int{
    []int{0,1},
    []int{0,-1},
    []int{1,0},
    []int{-1,0},
}
func dfs(matrix [][]int,i,j,m,n int,cache [][]int) int {
    if cache[i][j] != 0{
        return cache[i][j]
    }
    max := 1
    for _,d := range dir {
        x,y := i+d[0],j+d[1]
        if x < 0 || x >=m||y<0||y>=n || matrix[i][j]>=matrix[x][y]{
            continue
        }
        l := 1+dfs(matrix,x,y,m,n,cache)
        if max < l {
            max = l
        }
    }
    cache[i][j]=max
    return max
}
func main() {
    fmt.Println(longestIncreasingPath([][]int{
        []int{3,4,5},
        []int{3,2,6},
        []int{2,2,1},
    }))
    
}
