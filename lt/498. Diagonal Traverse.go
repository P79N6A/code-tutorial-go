package main

import "fmt"

/*

Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

Example:
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output:  [1,2,4,7,5,3,6,8,9]
Explanation:

Note:
The total number of elements of the given matrix will not exceed 10,000.
*/
func printup(matrix [][]int,i,j,m,n int) []int {
    res := make([]int,0)
    for ii,jj:=i,j;ii>=0&&jj<n;ii,jj = ii-1,jj+1 {
        res = append(res,matrix[ii][jj])
    }
    return res
}

func reverse(res []int) []int {
    for i:=0;i<len(res)/2;i++ {
        res[i],res[len(res)-1-i]= res[len(res)-1-i],res[i]
    }
    return res
}
func printdown(matrix [][]int,i,j,m,n int) []int {
    res := make([]int,0)
    for ii,jj:=i,j;ii<m&&jj>=0;ii,jj = ii+1,jj-1 {
        res = append(res,matrix[ii][jj])
    }
    return res
}
func findDiagonalOrder(matrix [][]int) []int {
    m := len(matrix)
    if m <= 0 {return nil}
    n := len(matrix[0])
    rr := make([]int,0)
    for i:=0;i<m;i+=1 {
        res := printup(matrix,i,0,m,n)
        if i%2 == 0 {
            rr = append(rr,res...)
        } else {
            rr = append(rr,reverse(res)...)
        }
    }
    for j:=1;j<n;j+=1 {
        res := printup(matrix,m-1,j,m,n)
        if (m-1+j)%2 == 0 {
            rr = append(rr,res...)
        } else {
            rr = append(rr,reverse(res)...)
        }
    }
    return rr
}

func main() {
    fmt.Println(findDiagonalOrder([][]int{
        []int{1},
    }))
    
}
