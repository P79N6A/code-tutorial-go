package main

import (
    "sort"
    "fmt"
)
/*
Given the coordinates of four points in 2D space, return whether the four points could construct a square.

The coordinate (x,y) of a point is represented by an integer array with two integers.

Example:
Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
Output: True
Note:

All the input integers are in the range [-10000, 10000].
A valid square has four equal sides with positive length and four equal angles (90-degree angles).
Input points have no order.
*/

func main() {
    fmt.Println(validSquare([]int{0,0},[]int{1,1},[]int{1,0},[]int{0,1}))
}
/*
先按照X轴排序,如果相等按照Y排序;这样0,1,2,3四个点的位置就固定了.
一共三种情况,分别判断边长,对角线是否相等即可.
*/
type points [][]int
func (p points)Len() int {
    return len(p)
}
func (p points)Swap(i,j int) {
    p[i],p[j]=p[j],p[i]
}
func (p points)Less(i,j int)bool {
    if p[i][0]==p[j][0] {
        return p[i][1] < p[j][1]
    }
    return p[i][0]<p[j][0]
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
    p := [][]int{p1,p2,p3,p4}
    sort.Sort(points(p))
    return dist(p[0], p[1]) != 0 &&  // 不重合
        dist(p[0],p[1])==dist(p[1],p[3]) &&
        dist(p[1],p[3])==dist(p[3],p[2]) &&
        dist(p[3],p[2])==dist(p[2],p[0]) &&  //边 0-1 == 1-3 == 3-2 == 2-0
        dist(p[0],p[3])==dist(p[1],p[2]) // 对角线
}
func dist(p1,p2 []int) int {
    return (p2[1] - p1[1]) * (p2[1] - p1[1]) + (p2[0] - p1[0]) * (p2[0] - p1[0])
}