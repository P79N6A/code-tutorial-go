package main

import "fmt"

/*
问题：给定的无向图，判断是否能将节点分成两组，使得所有相邻的节点都在不同的组中。

思路：经典的graph bipartite问题。
解题思路类似图染色问题，只不过现在只有两种颜色。
还是递归回溯，回溯过程中判断是否能够成功分组
如果是多个连通图，那就使用各个节点都判断一次

例题 890. Possible Bipartition
给出节点个数和一堆边，相邻的节点不能放在一组,问是否能够分成两组
并且说节点个数是有限的N<=2000
Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return true if and only if it is possible to split everyone into two groups in this way.



Example 1:

Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
Example 2:

Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Example 3:

Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false


Note:

1 <= N <= 2000
0 <= dislikes.length <= 10000
1 <= dislikes[i][j] <= N
dislikes[i][0] < dislikes[i][1]
There does not exist i != j for which dislikes[i] == dislikes[j].

参考 https://www.geeksforgeeks.org/bipartite-graph/
https://www.geeksforgeeks.org/check-if-a-given-graph-is-bipartite-using-dfs/
*/

func main() {
    fmt.Println(possibleBipartition(4,[][]int{
        {1,2},{1,3},{2,4},
    }))

}
func possibleBipartition(N int, dislikes [][]int) bool {
    adj := make([][]int,N+1)
    // 构建邻接表
    for _,dis := range dislikes {
        if len(adj[dis[0]]) <= 0 {adj[dis[0]]=make([]int,0)}
        if len(adj[dis[1]]) <= 0 {adj[dis[1]]=make([]int,0)}
        adj[dis[0]]=append(adj[dis[0]],dis[1])
        adj[dis[1]]=append(adj[dis[1]],dis[0])
    }
    // dfs过程，需要有visit控制是否访问， color记录访问节点的颜色
    visit := make([]bool,N+1)
    color := make([]bool,N+1)
    for i:=1;i<=N;i++ {
        // 可能是多个连通图，所以各个节点都做判断,多个连通图也可以分成2组
        // 如果已知是连通图，则直接visit[1]=true,color[1]=true 继续即可
        if dfs(adj,i,&visit,&color) == false {
            return false
        }
    }
    return true
}

func dfs(adj [][]int,v int,visit,color *[]bool) bool {
    /*
    搜索过程附带染色，如果发现已经染色子节点并且颜色冲突，则直接false
    如果有未染色节点，染上跟当前节点相反的颜色，继续下去
    */
    // input  graph, nodenum, currentidx, color map
    for _,u := range adj[v] {
        if (*visit)[u]==false {
            //如果没有访问过,标记访问状态和color状态
            (*visit)[u]=true
            (*color)[u]= !(*color)[v]
            if !dfs(adj,u,visit,color) {
                //只要错了就失败了
                return false
            }
        } else {
            if (*color)[u]==(*color)[v] {
                //如果遇到已访问节点，并且颜色和v一样，则肯定失败了
                return false
            }
        }
    }
    return true
}