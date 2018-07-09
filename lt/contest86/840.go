package main

import "fmt"
/*
A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

Given an N x N grid of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).



Example 1:

Input: [[4,3,8,4],
        [9,5,1,9],
        [2,7,6,2]]
Output: 1
Explanation:
The following subgrid is a 3 x 3 magic square:
438
951
276

while this one is not:
384
519
762

In total, there is only one magic square inside the given grid.
Note:

1 <= grid.length = grid[0].length <= 10
0 <= grid[i][j] <= 15

*/
func numMagicSquaresInside(grid [][]int) int {
        line := len(grid)
        row := len(grid[0])
        ret := 0
        for i:=0;i<=line-3;i++ {
                for j:=0;j<=row-3;j++ {
                        if is19(i,j,grid) {
                                ret += 1
                        }

                }
        }
        return ret
}
func is19(x,y int,grid [][]int) bool {
        if grid[x+1][y+1] != 5 {return false}
        for m:=x;m<x+3;m++ {
                sum1 := 0
                for n:=y;n<y+3;n++ {
                        if grid[m][n] > 9 {
                                return false
                        }
                        sum1 += grid[m][n]
                }
                if sum1 != 15 {
                        return false
                }
        }
        if grid[x][y]+grid[x+1][y]+grid[x+2][y] != 15 {return false}
        if grid[x][y+1]+grid[x+1][y+1]+grid[x+2][y+1] != 15 {return false}
        if grid[x][y+2]+grid[x+1][y+2]+grid[x+2][y+2] != 15 {return false}

        if grid[x][y]+grid[x+1][y+1]+grid[x+2][y+2] != 15 {return false}
        if grid[x+2][y]+grid[x+1][y+1]+grid[x][y+2] != 15 {return false}
        return true
}
func main() {
        fmt.Println(numMagicSquaresInside([][]int{
                []int{3,2,9,2,7},
                []int{6,1,8,4,2},
                []int{7,5,3,2,7},
                []int{2,9,4,9,6},
                []int{4,3,8,2,5},
        }))
}
