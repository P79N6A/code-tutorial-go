package main

import "fmt"

/*
An undirected, connected graph of N nodes (labeled 0, 1, 2, ..., N-1) is given as graph.

graph.length = N, and j != i is in the list graph[i] exactly once, if and only if nodes i and j are connected.

Return the length of the shortest path that visits every node. You may start and stop at any node, you may revisit nodes multiple times, and you may reuse edges.



Example 1:

Input: [[1,2,3],[0],[0],[0]]
Output: 4
Explanation: One possible path is [1,0,2,0,3]
Example 2:

Input: [[1],[0,2,4],[1,3,4],[2],[1,2]]
Output: 4
Explanation: One possible path is [0,1,4,2,3]


Note:

1 <= graph.length <= 12
0 <= graph[i].length < graph.length
*/

type State struct {
    cur int
    head int
}
func shortestPathLength(graph [][]int) int {
    // key is state. value is distance
    //这个dist两个用途，第一个用来记录最小的distance，第二个就是用来去重queue中的state
    // 使用一个针对cur的map也是可以的，因为第一次进来的总是最小的
    dist := make([][]int,0)
    for i:=0;i<len(graph);i++ {dist = append(dist,make([]int,1<<uint(len(graph))))}
    queue := make([]State,0)
    for i:=0;i<len(graph);i++ {
        // 初始化，将所有的顶点都放进去，state，distance
        queue = append(queue,State{1<<uint(i),i})
        dist[i][1<<uint(i)]=0
    }
    for len(queue) > 0 {
        // 获得队列头的state和distance
        //根据state的head，看下一批加入的节点有哪些
        t := queue[0]
        d := dist[t.head][t.cur]
        if t.cur == (1<<uint(len(graph))-1) {
            return d
        }
        for _,e := range graph[t.head] {
            ncur := t.cur | 1<<uint(e)
            if dist[e][ncur] == 0 || dist[e][ncur] > d+1 {
                // 用来去重复的。
                dist[e][ncur]=d+1
                queue=append(queue,State{ncur,e})
            }
        }
        queue=queue[1:]

    }
    return -1
}
func main() {
    fmt.Println(shortestPathLength([][]int{
        []int{1,2,3},
        []int{0},
        []int{0},
        []int{0},
    }))
    
}
