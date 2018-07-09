package main

import (
        "fmt"
)
/*

Given a triangle, find the minimum path sum from top to bottom.
Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.


*/
func minimumTotal(triangle [][]int) int {
        line := len(triangle)
        if line <= 0 {return 0}
        row := len(triangle[line-1])
        dp := make([]int,row+1)
        for i:=line-1;i>=0;i-- {
                for j:=0;j<=i;j++ {
                        if dp[j]>dp[j+1] {
                                dp[j]=dp[j+1]+triangle[i][j]
                        } else {
                                dp[j]=dp[j]+triangle[i][j]
                        }
                }
        }
        return dp[0]
}

func main() {
        fmt.Println(minimumTotal([][]int{
                []int{2},
                []int{3,4},
                []int{6,5,7},
                []int{4,1,8,3},
        }))
}
