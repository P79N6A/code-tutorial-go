package main

import "fmt"

/*
On a 2D plane, we place stones at some integer coordinate points.  Each coordinate point may have at most one stone.

Now, a move consists of removing a stone that shares a column or row with another stone on the grid.

What is the largest possible number of moves we can make?



Example 1:

Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
Output: 5
Example 2:

Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
Output: 3
Example 3:

Input: stones = [[0,0]]
Output: 0


Note:

1 <= stones.length <= 1000
0 <= stones[i][j] < 10000
*/

func main() {
    fmt.Println(removeStones([][]int{{0,0},{0,1},{1,0},{1,2},{2,1},{2,2}}))
    fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
    fmt.Println(removeStones([][]int{{0,0}}))
}

/*
问题的意思是，将所有的行或者列相同的点都放到一组里面。问到底能有多少分组？

解法1：DFS，遍历；看有多少个island。如何表达图关系？
解法2：Union Find，并查集
*/
func removeStones(stones [][]int) int {
    rowg := make(map[int][][2]int)
    lineg := make(map[int][][2]int)
    for _, s := range stones { // 构建邻接表
        if len(lineg[s[0]]) <= 0 {
            lineg[s[0]] = make([][2]int, 0)
        }
        lineg[s[0]] = append(lineg[s[0]], [2]int{s[0], s[1]})
        if len(rowg[s[1]]) <= 0 {
            rowg[s[1]] = make([][2]int, 0)
        }
        rowg[s[1]] = append(rowg[s[1]], [2]int{s[0], s[1]})
    }
    group := 0  //  按分了多少组
    visit := make(map[[2]int]bool)
    for _, s := range stones {
        n := [2]int{s[0], s[1]}
        if !visit[n] {
            dfs(n, visit, rowg, lineg)
            group += 1
        }
    }
    return len(stones) - group
}
func dfs(node [2]int, visit map[[2]int]bool, rows, lines map[int][][2]int) {
    visit[node] = true
    for _, r := range rows[node[1]] {
        if !visit[r] {
            dfs(r, visit, rows, lines)
        }
    }
    for _, l := range lines[node[0]] {
        if !visit[l] {
            dfs(l, visit, rows, lines)
        }
    }
}
///////////////////////