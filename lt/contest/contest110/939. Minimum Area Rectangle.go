package main

import (
    "sort"
    "math"
    "fmt"
)
/*
Given a set of points in the xy-plane, determine the minimum area of a rectangle formed from these points, with sides parallel to the x and y axes.

If there isn't any rectangle, return 0.



Example 1:

Input: [[1,1],[1,3],[3,1],[3,3],[2,2]]
Output: 4
Example 2:

Input: [[1,1],[1,3],[3,1],[3,3],[4,1],[4,3]]
Output: 2


Note:

1 <= points.length <= 500
0 <= points[i][0] <= 40000
0 <= points[i][1] <= 40000
All points are distinct.

*/

func main() {
    //fmt.Println(minAreaRect([][]int{{1,1},{1,3},{3,1},{3,3},{2,2}}))
    fmt.Println(minAreaRect([][]int{{1,1},{1,3},{3,1},{3,3},{4,1},{4,3}}))
}
// 注意题目只是要求跟x,y平行的矩形,,so,想多了,非平行矩形不在考虑范围.
func minAreaRect(points [][]int) int {
    xmap := make(map[int]map[int]bool)//相同的x放到一个set里面
    for _,p := range points {
        if len(xmap[p[0]]) <= 0 {
            xmap[p[0]]=make(map[int]bool)
        }
        xmap[p[0]][p[1]]=true
    }
    ans := math.MaxInt64
    for i:=0;i<len(points);i++ {
        for j:=i+1;j<len(points);j++ {
            //任意的在一条X平行线或者Y平行线的两点是无法构成矩形的
            if points[i][0]==points[j][0] || points[i][1]==points[j][1] {
                continue
            }
            //判断对称的坐标是否在对方集合中,points[i],points[j]是对角线的两点.
            if xmap[points[i][0]][points[j][1]]&&xmap[points[j][0]][points[i][1]] {
                rect := (points[i][0]-points[j][0])*(points[i][1]-points[j][1])
                if rect < 0 {rect*=-1} // math.abs
                if rect>0&&rect<ans{ans=rect} // math.min
            }
        }
    }
    if ans == math.MaxInt64 {return 0}
    return ans
}

func dist(p1,p2 []int) int {
    return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[1]-p1[1])*(p2[1]-p1[1])
}
func check(p1,p2,p3,p4 []int) bool {
    return dist(p1,p2)>0&&dist(p1,p2)==dist(p2,p3)&&dist(p2, p3) == dist(p3, p4) && dist(p3, p4) == dist(p4, p1) && dist(p1, p3) == dist(p2, p4)
}
/*
    public boolean validSquare(int[] p1, int[] p2, int[] p3, int[] p4) {
        int[][] p={p1,p2,p3,p4};
        Arrays.sort(p, (l1, l2) -> l2[0] == l1[0] ? l1[1] - l2[1] : l1[0] - l2[0]);
        return dist(p[0], p[1]) != 0 && dist(p[0], p[1]) == dist(p[1], p[3]) && dist(p[1], p[3]) == dist(p[3], p[2]) && dist(p[3], p[2]) == dist(p[2], p[0])   && dist(p[0],p[3])==dist(p[1],p[2]);
    }
}
*/
func isSquare(p1,p2,p3,p4 []int) (bool,int) {
    p := [][]int{p1,p2,p3,p4}
    sort.Sort(squarePoint(p))
    if p[0][0]==p[1][0] && p[0][1]!=p[1][1]{
        if p[0][1]==p[2][1]&&p[1][1]==p[3][1]&&p[2][0]==p[3][0] {
            fmt.Println(p)
            return true,(p[1][1]-p[0][1])*(p[2][0]-p[0][0])
        }
    } else {
        if dist(p[0], p[1]) != 0 && dist(p[0], p[1]) == dist(p[2], p[3])&&dist(p[0], p[2]) == dist(p[1], p[3])&&dist(p[0], p[3]) == dist(p[1], p[2])&&dist(p[0], p[3]) == dist(p[0], p[2]) + dist(p[2], p[3]) {
            return true,int(math.Sqrt(float64(dist(p[0],p[1])*dist(p[0],p[2]))))
        }
    }
    return false,-1
}
type squarePoint [][]int
func (l squarePoint)Len() int {
    return len(l)
}
func (l squarePoint)Swap(i,j int) {
    l[i],l[j]=l[j],l[i]
}
func (l squarePoint)Less(i,j int) bool {
    if l[i][0]==l[j][0] {
        return l[i][1]<=l[j][1]
    }
    return l[i][0]<l[j][0]
}
/*
https://leetcode.com/problems/valid-square/solution/
    public double dist(int[] p1, int[] p2) {
        return (p2[1] - p1[1]) * (p2[1] - p1[1]) + (p2[0] - p1[0]) * (p2[0] - p1[0]);
    }
    public boolean check(int[] p1, int[] p2, int[] p3, int[] p4) {
        return dist(p1,p2) > 0 && dist(p1, p2) == dist(p2, p3) && dist(p2, p3) == dist(p3, p4) && dist(p3, p4) == dist(p4, p1) && dist(p1, p3) == dist(p2, p4);
    }
    public boolean validSquare(int[] p1, int[] p2, int[] p3, int[] p4) {
        return check(p1, p2, p3, p4) || check(p1, p3, p2, p4) || check(p1, p2, p4, p3);
    }
*/
