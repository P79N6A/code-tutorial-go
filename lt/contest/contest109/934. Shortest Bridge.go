package main
/*
In a given 2D binary array A, there are two islands.  (An island is a 4-directionally connected group of 1s not connected to any other 1s.)

Now, we may change 0s to 1s so as to connect the two islands together to form 1 island.

Return the smallest number of 0s that must be flipped.  (It is guaranteed that the answer is at least 1.)



Example 1:

Input: [[0,1],[1,0]]
Output: 1
Example 2:

Input: [[0,1,0],[0,0,0],[0,0,1]]
Output: 2
Example 3:

Input: [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
Output: 1


Note:

1 <= A.length = A[0].length <= 100
A[i][j] == 0 or A[i][j] == 1
*/

import "fmt"

func main() {
    fmt.Println(shortestBridge([][]int{{1,1,1,1,1},{1,0,0,0,1},{1,0,1,0,1},{1,0,0,0,1},{1,1,1,1,1}}))//1
    fmt.Println(shortestBridge([][]int{{0,1,0},{0,0,0},{0,0,1}}))//2
    fmt.Println(shortestBridge([][]int{{0,1},{1,0}}))//1
}
/*
问题:0,1组成的二维数组构成一个图,1代表的岛屿,题目中说只有两个岛.问最少需要flip翻转几个0能将两个岛屿连在一起
思路:非常非常经典的图的遍历.使用DFS做染色. 两个岛染上不同的颜色. 使用BFS计算距离.
比如最终染色得到B,C两个岛,遍历节点个数少的,比如B,看B上任意一点到C的距离的最小值
BFS计算路径的方法是每次都统一遍历完queue中所有元素
*/
func shortestBridge(A [][]int) int {
    if len(A)<=0 {return 0}
    color := 2
    nums := make(map[int]int)
    for i:=0;i<len(A);i++ {
        for j:=0;j<len(A[i]);j++ {
            if A[i][j]== 1 {
                num := 0
                dfs(i, j, color, &A, &num, len(A), len(A[0]))
                if num > 0 {
                    nums[color] = num
                    color += 1
                }
            }
        }
    }
    isl,target := 2,3
    if nums[2]>nums[3] {
        isl,target = 3,2
    }
    ans := len(A)*len(A[0])
    for i:=0;i<len(A);i++ {
        for j := 0; j < len(A[i]); j++ {
            if A[i][j]==isl {
                x := bfs(i,j,target,A,len(A),len(A[0]))
                if x < ans {
                    ans = x
                }
            }
        }
    }
    return ans
}
func bfs(i,j int,target int,A[][]int,m,n int) int {
    visit := make([][]int,0)
    for i:=0;i<len(A);i++ {
        visit = append(visit,make([]int,len(A[i])))
    }
    dis := 0
    queue := make([][]int,0)
    queue = append(queue,[]int{i,j})
    visit[i][j]=1
    for len(queue)>0 {
        seq := len(queue)
        for i:=0;i<seq;i++ { // 每次遍历完一次相同距离的所有节点
            t := queue[i]
            if A[t[0]][t[1]]==target {
                return dis-1
            }
            for _,d := range [][]int{{0,1},{1,0},{0,-1},{-1,0}} {
                x,y := d[0]+t[0],d[1]+t[1]
                if x<0||x>=m||y<0||y>=n||visit[x][y]==1{
                    continue
                }
                visit[x][y]=1
                queue = append(queue,[]int{x,y})
            }
        }
        dis += 1 //  距离
        queue = queue[seq:] // 遍历完相同距离所有节点后后移
    }
    return dis
}

func dfs(i,j int, color int, A *[][]int, num *int,m,n int) {
    // 染色, num染色节点个数, color目标颜色
    (*A)[i][j]=color
    *num += 1
    for _,d := range [][]int{{0,1},{1,0},{0,-1},{-1,0}} {
        x,y := d[0]+i,d[1]+j
        if x<0||x>=m||y<0||y>=n||(*A)[x][y]!=1{
            continue
        }
        dfs(x,y,color,A,num,m,n)
    }
}