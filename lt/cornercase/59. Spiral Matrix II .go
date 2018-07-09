package cornercase

import "fmt"

/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/
func generateMatrix(n int) [][]int {
    matrix := make([][]int,0)
    for i:=0;i<n;i++ {
        matrix = append(matrix,make([]int,n))
    }
    sx,sy,row,col,t := 0,0,n,n,1
    for sx < row && sy < col {
        t = generate(matrix,t,sx,sy,row,col)
        sx += 1
        sy += 1
        row -= 1
        col -= 1
    }
    return matrix
}
func generate(matrix [][]int,target int, x,y,ex,ey int) int {
    for i:=y;i<ey;i++ {matrix[x][i]=target;target += 1}
    for i:=x+1;i<ex;i++ {matrix[i][ey-1]=target;target += 1}
    for i:=ey-2;i>=y&&x+1<ex;i-- {matrix[ex-1][i]=target;target+=1}
    for i:=ex-2;i>=x+1&&ey-1>y;i-- {matrix[i][y]=target;target+=1}
    return target
}

func main() {
    m := generateMatrix(3)
    for _,mm := range m {
        fmt.Println(mm)
    }
}
