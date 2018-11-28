package main

import "fmt"

/*
In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input:
nums =
[[1,2],
 [3,4]]
r = 1, c = 4
Output:
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
Example 2:
Input:
nums =
[[1,2],
 [3,4]]
r = 2, c = 4
Output:
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
Note:
The height and width of the given matrix is in range [1, 100].
The given r and c are all positive.
*/

func main() {
    fmt.Println(matrixReshape([][]int{{1,2},{3,4}},1,4))
    fmt.Println(matrixReshape([][]int{{1},{2},{3},{4}},2,2))

}
func matrixReshape(nums [][]int, r int, c int) [][]int {
    // 额外空间，使用queue
    if len(nums)<=0 {return nums}
    m,n := len(nums),len(nums[0])
    if m*n != r*c {return nums}
    queue := make([]int,0,r*c)
    for i:=0;i<m;i++ {
        for j := 0; j < n; j++ {
            queue = append(queue,nums[i][j])
        }
    }
    ans := make([][]int,0)
    for i:=0;i<r;i++ {
        ans = append(ans,make([]int,c))
    }
    idx := 0
    for i:=0;i<r;i++ {
        for j:=0;j<c;j++ {
            ans[i][j]=queue[idx]
            idx += 1
        }
    }
    return ans
}
func matrixReshape1(nums [][]int, r int, c int) [][]int {
    /*
    使用除法
    */
    if len(nums)<=0 {return nums}
    m,n := len(nums),len(nums[0])
    if m*n != r*c {return nums}
    ans := make([][]int,0)
    for i:=0;i<r;i++ {
        ans = append(ans,make([]int,c))
    }
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            ii := (i*n+j)/c
            jj := (i*n+j)%c
            ans[ii][jj]=nums[i][j]
        }
    }
    return ans
}
