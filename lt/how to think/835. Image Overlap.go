package main

import "fmt"

/*
Two images A and B are given, represented as binary, square matrices of the same size.  (A binary matrix has only 0s and 1s as values.)

We translate one image however we choose (sliding it left, right, up, or down any number of units), and place it on top of the other image.  After, the overlap of this translation is the number of positions that have a 1 in both images.

(Note also that a translation does not include any kind of rotation.)

What is the largest possible overlap?

Example 1:

Input: A = [[1,1,0],
            [0,1,0],
            [0,1,0]]
       B = [[0,0,0],
            [0,1,1],
            [0,0,1]]
Output: 3
Explanation: We slide A to right by 1 unit and down by 1 unit.
Notes:

1 <= A.length = A[0].length = B.length = B[0].length <= 30
0 <= A[i][j], B[i][j] <= 1
*/
/*是平移，不是翻转。。。 复杂度上升了N^2*/
func main() {
    fmt.Println(largestOverlap([][]int{
        {1,1,0},
        {0,1,0},
        {0,1,0},
    },[][]int{
        {0,0,0},
        {0,1,1},
        {0,0,1},
    }))
    fmt.Println(largestOverlap([][]int{
        {0,1},
        {1,1},
    },[][]int{
        {1,1},
        {1,0},
    }))
    fmt.Println(largestOverlap([][]int{
        {1,0},
        {1,0},
    },[][]int{
        {1,1},
        {1,0},
    })) // 2

}
/////////////////////////////
func largestOverlap(A [][]int, B [][]int) int {
    if len(A) <= 0 {return 0}
    n := len(A[0])
    aVector := make([]int,0)
    for i:=0;i<n;i++ {
        for j:=0;j<n;j++ {
            // A中1的坐标，只不过坐标没用二维向量，而是使用i*100+j
            if A[i][j]==1 {aVector = append(aVector,i*100+j)}
        }
    }
    bVector := make([]int,0)
    for i:=0;i<n;i++ {
        for j:=0;j<n;j++ {
            // B中1的坐标，只不过坐标没用二维向量，而是使用i*100+j
            if B[i][j]==1 {bVector = append(bVector,i*100+j)}
        }
    }
    counter := make(map[int]int)
    for _,a := range aVector {
        for _,b := range bVector {
            //计算A,B中所有的1做向量平移
            // 向量作为Key, value是个数，最终得到的相同的向量平移一共多少个点
            // a-b ==> (ax-bx,ay-by)
            counter[a-b] += 1
        }
    }
    max := 0
    for _,c := range counter {if c > max {max = c}}
    return max
}
/////////////////////////////
/*
两个只有0,1的二维数组A,B，矩形长度N，可以做上下左右的平移操作。平移后叠放在一起，1的个数加和即为Overlap，问Overlap最大值。
思路：如果是rotate就简单的多，固定B,遍历A的i,j作为新的起点，看跟B重合的1的个数即可。。
比较直接的想法：画个二维数轴，把两个矩形放进去， A的左上和右下坐标(0,0),(n,-n)，这样A直接遍历数组不必坐标变换;那么B的左下坐标范围是应该是(-N,-N)=>(N,N) 一共2N*2N个点，那么问题转化为遍历任意一种可能，计算A,B相交部分的1的个数。这种计算涉及坐标系转换。
*/

func largestOverlap1(A [][]int, B [][]int) int {
    if len(A) <= 0 {return 0}
    n := len(A[0])
    max := 0
    // A在坐标位置就是左上(0,0)，右下(n,-n) 方便遍历数组
    // B左上坐标在范围 (-n,n)=>(n,-n)；需要坐标转换
    for i:=-n;i<=n;i++ {
        for j:=-n;j<=n;j++ {
            c := cal(A, B, i, j, n)
            if c > max {max = c}
        }
    }
    return max
}
func cal(A,B [][]int,ax,ay,n int) int {
    /*
    固定A，通过坐标转换[绝对坐标转换成相对的]B的(0,0)->(n,n)坐标
    */
    num := 0
    for i:=0;i<n;i++ {
        for j:=0;j<n;j++ {
            if A[i][j]==0{
                continue
            }
            // 坐标转换，向量减法
            bi,bj := i-ax,j-ay
            // 转换后在有效范围内
            if bi >= 0 && bi <n && bj >=0&&bj<n {
                if B[bi][bj]==1 {num +=1}
            }
        }
    }
    return num
}

