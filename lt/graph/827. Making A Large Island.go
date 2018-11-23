package main

import "fmt"

/*
In a 2D grid of 0s and 1s, we change at most one 0 to a 1.

After, what is the size of the largest island? (An island is a 4-directionally connected group of 1s).

Example 1:

Input: [[1, 0], [0, 1]]
Output: 3
Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.
Example 2:

Input: [[1, 1], [1, 0]]
Output: 4
Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.
Example 3:

Input: [[1, 1], [1, 1]]
Output: 4
Explanation: Can't change any 0 to 1, only one island with area = 4.


Notes:

1 <= grid.length = grid[0].length <= 50.
0 <= grid[i][j] <= 1.

*/

func main() {
    fmt.Println(largestIsland([][]int{{1, 0}, {0, 1}}))
    fmt.Println(largestIsland([][]int{{1, 1}, {1, 0}}))
    fmt.Println(largestIsland([][]int{{1, 1}, {1, 1}}))
    x := [][]int{
        {0,0,0,0,0,0,0},{0,1,1,1,1,0,0},{0,1,0,0,1,0,0},{1,0,1,0,1,0,0},{0,1,0,0,1,0,0},{0,1,0,0,1,0,0},{0,1,1,1,1,0,0},
    }
    fmt.Println(largestIsland(x))
}
func largestIsland(grid [][]int) int {
    if len(grid) <= 0 {return 0}
    dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    m, n := len(grid), len(grid[0])
    /*
    第一个返回值表示当前最大island的大小
    第二个返回值表示island分组
    并且处理grid，将同一个islan的grid值修改成island大小
    */
    ans,group := largestIs(&grid, m, n)
    ans += 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j]==0 {
                /*
                找所有是0的点，看0的点连接了几个group
                */
                one := 1 // 连接island大小，默认是1，因为自己变成1
                gset := make(map[int]bool) // group不能重复
                for _, d := range dir {
                    u,v := i+d[0],j+d[1]
                    if u<0||u>=m||v<0||v>=n||grid[u][v]==0 {continue}
                    if _,ok := gset[group[[2]int{u,v}]];!ok {
                        gset[group[[2]int{u,v}]]=true
                        one += grid[u][v] // 将每个独立group的island大小加和
                    }
                }
                if one > ans {ans = one}
            }
        }
    }
    if ans > m*n {
        return m * n
    }
    return ans
}
func largestIs(grid *[][]int, m, n int) (int,map[[2]int]int) {
    ans := 0
    visit := make([][]bool, 0)
    for i := 0; i < m; i++ {
        visit = append(visit, make([]bool, n))
    }
    group := make(map[[2]int]int) // key是坐标点，value是groupid

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !visit[i][j] && (*grid)[i][j] == 1 {
                visit[i][j] = true
                nodes := make([][]int,0)
                nodes = append(nodes,[]int{i,j})
                one := 1
                dfs(i, j, grid, &visit, m, n, &one, &nodes)
                for _,node := range nodes {
                    (*grid)[node[0]][node[1]]=one
                    group[[2]int{node[0],node[1]}]=i*m+j
                }
                if one > ans {
                    ans = one
                }
            }
        }
    }
    return ans,group
}
func dfs(i, j int, grid *[][]int, visit *[][]bool, m, n int, ans *int,nodes *[][]int) {
    // ans代表包含i,j的island的大小
    // nodes表示这个island的坐标点的集合
    dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    for _, d := range dir {
        u, v := i+d[0], j+d[1]
        if u < 0 || u >= m || v < 0 || v >= n || (*visit)[u][v] == true || (*grid)[u][v] == 0 {
            continue
        }
        *ans += 1
        *nodes = append(*nodes,[]int{u,v})
        (*visit)[u][v] = true
        dfs(u, v, grid, visit, m, n, ans,nodes)
    }
}
