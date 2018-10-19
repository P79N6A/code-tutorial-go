package main

import "fmt"

/*
Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.

The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j for which the edge (i, j) exists.

Example:
Input: [[1,2], [3], [3], []]
Output: [[0,1,3],[0,2,3]]
Explanation: The graph looks like this:
0--->1
|    |
v    v
2--->3
There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
Note:

The number of nodes in the graph will be in the range [2, 15].
You can print different paths in any order, but you should keep the order of nodes inside one path.
*/

func main() {
    fmt.Println(allPathsSourceTarget([][]int{
        {1,2},{3},{3},{},
    }))
}
func allPathsSourceTarget(graph [][]int) [][]int {
    ans := make([][]int,0)
    once := make([]int,0)
    visit := make([]int,len(graph))
    visit[0]=1
    once = append(once,0)
    dfs(0,graph,len(graph)-1,&visit,&once,&ans)
    return ans
}
func dfs(u int, g [][]int, t int,visit,once *[]int, ans *[][]int) {
    if u == t {
        r := make([]int,len(*once))
        copy(r,*once)
        *ans = append(*ans,r)
        return
    }
    for _,v := range g[u] {
        if (*visit)[v]==0 {
            (*visit)[v]=1
            *once = append(*once,v)
            dfs(v,g,t,visit,once,ans)
            (*visit)[v]=0
            *once=(*once)[:(len(*once)-1)]
        }
    }
}