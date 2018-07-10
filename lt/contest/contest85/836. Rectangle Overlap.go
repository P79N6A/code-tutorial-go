package main

import "fmt"
/*
A rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) are the coordinates of its bottom-left corner,
and (x2, y2) are the coordinates of its top-right corner.

Two rectangles overlap if the area of their intersection is positive.
To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two rectangles, return whether they overlap.

Example 1:

Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
Output: true
Example 2:

Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
Output: false
Notes:

Both rectangles rec1 and rec2 are lists of 4 integers.
All coordinates in rectangles will be between -10^9 and 10^9.
*/
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
        return judgeLineCross(rec1[0],rec1[2],rec2[0],rec2[2]) && judgeLineCross(rec1[1],rec1[3],rec2[1],rec2[3])
}
func judgeLineCross(x,y,m,n int) bool {
        fmt.Println(x,y,m,n)
        if (x <=m && y <= m) || (x>=n&&y>=n) {return false}
        return true
}
func main() {
        //fmt.Println(isRectangleOverlap([]int{0,0,2,2},[]int{1,1,3,3}))
        //fmt.Println(isRectangleOverlap([]int{0,0,1,1},[]int{1,0,2,1}))
        //fmt.Println(isRectangleOverlap([]int{1,1,1,1},[]int{1,0,2,1}))
        fmt.Println(isRectangleOverlap([]int{4,0,6,6},[]int{-5,-3,4,2}))
}
