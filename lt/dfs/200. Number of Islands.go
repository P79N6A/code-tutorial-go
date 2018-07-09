package main

import "fmt"
/*
Given a 2d grid map of '1's (land) and '0's (water),
count the number of islands. An island is surrounded
by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3

*/
func numIslands(grid [][]byte) int {
        line := len(grid)
        if line < 1 {return 0}
        row := len(grid[0])
        counter := 0
        for i:=0;i<line;i++ {
                for j:=0;j<row;j++ {
                        if grid[i][j]=='1' {
                                dfs(grid,line,row,i,j)
                                counter += 1
                        }
                }
        }
        return counter
}
func dfs(grid [][]byte,line,row,i,j int) {
        if i < 0 || i >= line || j < 0 || j >= row {
                return
        }
        if grid[i][j] == '0' {
                return
        }
        grid[i][j]='0'
        dfs(grid,line,row,i+1,j)
        dfs(grid,line,row,i-1,j)
        dfs(grid,line,row,i,j+1)
        dfs(grid,line,row,i,j-1)

}

func main() {
        fmt.Println(numIslands([][]byte{
                []byte("11110"),
                []byte("11010"),
                []byte("11000"),
                []byte("00000"),
        }))
}
