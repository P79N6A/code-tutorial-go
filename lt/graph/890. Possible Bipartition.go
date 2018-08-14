package main

import "fmt"

/*
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
*/

func main() {
    fmt.Println(possibleBipartition(4,[][]int{
        {1,2},{1,3},{2,4},
    }))
}
func possibleBipartition(N int, dislikes [][]int) bool {
    adj := make([][]int,N+1) // 构建邻接表
    for _,dis := range dislikes {
        if len(adj[dis[0]]) <= 0 {adj[dis[0]]=make([]int,0)}
        if len(adj[dis[1]]) <= 0 {adj[dis[1]]=make([]int,0)}
        adj[dis[0]]=append(adj[dis[0]],dis[1])
        adj[dis[1]]=append(adj[dis[1]],dis[0])
    }
    color := make([]int,N+1)
    color[1]=1 // 初始颜色
    for i:=1;i<=N;i++ {
        // 逐个进行染色，验证，因为染色过程判断了是否已经染色，所以不会重复染色
        dfs(adj,i,&color) // 染色
        if valid(adj,i,&color) == false { //验证
            return false
        }
    }
    return true
}
func valid(adj [][]int,v int, color *[]int) bool {
    // 对节点v的染色结果进行验证，看是否有其他的边跟他同色
    for _,u := range adj[v] {
        if (*color)[v] == (*color)[u] {
            return false
        }
    }
    return true
}
func dfs(adj [][]int,v int, color *[]int) {
    /*
    只管染色，相邻的没有染色的，染成相反的颜色
    */
    for _,u := range adj[v] {
        if (*color)[u] == 0 {  // 只处理一次，处理过的不管了。留着验证就行了
            (*color)[u]=-(*color)[v] // 反色
            dfs(adj,u,color)
        }
    }
}