package cornercase

import "fmt"

/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func spiralOrder(matrix [][]int) []int {
    row := len(matrix)
    if row <= 0 {return []int{}}
    col := len(matrix[0])
    sx,sy := 0,0
    res := make([]int,0)
    for sx < row && sy < col {
        r := spiral(matrix,sx,sy,row,col)
        res = append(res,r...)
        sx += 1
        sy += 1
        row -= 1
        col -= 1
    }
    return res
}
func spiral(matrix [][]int, x,y,ex,ey int) []int {
    res := make([]int,0)
    for i:=y;i<ey;i++ {res = append(res,matrix[x][i])}
    for i:=x+1;i<ex;i++ {res = append(res,matrix[i][ey-1])}
    for i:=ey-2;i>=y&&x+1<ex;i-- {res = append(res,matrix[ex-1][i])}
    for i:=ex-2;i>=x+1&&ey-1>y;i-- {res = append(res,matrix[i][y])}
    return res
}

func main() {
    matrix := [][]int{
        []int{1,2,3},
        []int{4,5,6},
        []int{7,8,9},
        //[]int{10,11,12},
        //[]int{13,14,15},
    }
    fmt.Println(spiralOrder(matrix))
    //fmt.Println(spiral(matrix,0,0,3,3))
    //fmt.Println(spiral(matrix,1,1,4,2))

    matrix1 := [][]int{
        []int{1,2,3,4},
        []int{5,6,7,8},
        []int{9,10,11,12},
        //[]int{10,11,12},
        //[]int{13,14,15},
    }
    //fmt.Println(spiralOrder(matrix1))
    fmt.Println(spiral(matrix1,1,1,2,3))
}
