package main

import "fmt"

/*
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
Example 2:
[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
*/

func main() {
    fmt.Println(maxAreaOfIsland([][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}})) // 4
}
func maxAreaOfIsland(grid [][]int) int {
    max := 0
    for i:=0;i<len(grid);i++ {
        for j:=0;j<len(grid[i]);j++ {
            if grid[i][j]==1 {
                x := dfs(grid, i, j)
                if x > max {max = x}
            }
        }
    }
    return max
}
var dirs = [][]int{{1,0},{0,1},{-1,0},{0,-1}} // 控制方向
func dfs(grid [][]int,i,j int) int {
    //不能重复调用。
    if i<0 || i >= len(grid) || j<0 || j >=len(grid[i]) || grid[i][j] != 1 {
        // 套路写法，判断有效性。
        return 0
    }
    grid[i][j]=-1 //  grid可以被修改，省去visit数组
    ret := 1 // 说明本身是1，有效结果，自己加个1
    for _,d := range dirs {
        ret += dfs(grid,i+d[0],j+d[1])
    }
    return ret
}
