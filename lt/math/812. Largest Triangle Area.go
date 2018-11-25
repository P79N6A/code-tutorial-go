package main

import "math"
/*
You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

Example:
Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
Output: 2
Explanation:
The five points are show in the figure below. The red triangle is the largest.


Notes:

3 <= points.length <= 50.
No points will be duplicated.
 -50 <= points[i][j] <= 50.
Answers within 10^-6 of the true value will be accepted as correct.

*/

func main() {
}
/*
给定三个坐标点,如何求三角形面积?
思路:将一个点通过向量平移转到(0,0)
we could derive it for triangles with the following approach:
 starting with points (px, py), (qx, qy), (rx, ry),
 the area of this triangle is the same under a translation by (-rx, -ry),
 so that the points become (px-rx, py-ry), (qx-rx, qy-ry), (0, 0).
*/
func largestTriangleArea(points [][]int) float64 {
    N := len(points)
    var ans float64

    for i:=0;i<N;i++ {
        for j:=i+1;j<N;j++ {
            for k:=j+1;k<N;k++ {
                ans = math.Max(ans,area(points[i],points[j],points[k]))
            }
        }
    }
    return ans
}
func area(p,q,r []int) float64 {
    return 0.5*math.Abs(float64(p[0]*q[1] + q[0]*r[1] + r[0]*p[1]-
                                p[1]*q[0] - q[1]*r[0] - r[1]*p[0]))
}
