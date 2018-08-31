package main

import "fmt"

/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

func main() {
    m := [][]int{
        {5, 1, 9,11},
        {2, 4, 8,10},
        {13, 3, 6,7},
        {15,14,12,16},
    }
    rotate(m)
    printm(m)
    m2 := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    rotate(m2)
    printm(m2)
    m3 := [][]int{{1}}
    rotate(m3)
    printm(m3)
}
func printm(m [][]int) {
    for _,mm := range m {
        fmt.Println(mm)
    }
}
func rotate(matrix [][]int)  {
    step := len(matrix)-1
    for i,j:=0,0;step>=1;i,j=i+1,j+1 {
        rotateOnce(matrix,i,j,step)
        step -=2
    }
}
func rotateOnce(m [][]int,si,sj int, step int) {
    for i:=0;i<step;i++ {
        tmp := m[si][sj+i]
        m[si][sj+i]=m[si+step-i][sj]
        m[si+step-i][sj]=m[si+step][sj+step-i]
        m[si+step][sj+step-i]=m[si+i][sj+step]
        m[si+i][sj+step]=tmp
    }
}