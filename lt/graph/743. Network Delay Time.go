package main

import (
    "math"
    "fmt"
)

/*
There are N network nodes, labelled 1 to N.

Given times, a list of travel times as directed edges times[i] = (u, v, w), where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.

Now, we send a signal from a certain node K. How long will it take for all nodes to receive the signal? If it is impossible, return -1.

Note:
N will be in the range [1, 100].
K will be in the range [1, N].
The length of times will be in the range [1, 6000].
All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 1 <= w <= 100.

*/

func main() {
    fmt.Println(networkDelayTime([][]int{{2,1,1},{2,3,1},{3,4,1}},4,2))
    
}
/*
问题：从K出发，在有相同中到所有的顶点距离的最大值[同时的网络传输,到所有节点的时间]
思路：dijkstra 算法；给定点到任意点的最短距离；
维护一个最短距离数组，每次从未处理的点中选出最短的，加入后，看是否能更新给定点到未处理点的最短距离
*/
type edge struct {
    v,w int
}
func networkDelayTime(times [][]int, N int, K int) int {
    graph := make([][]edge,N+1) // 构建邻接表
    for _,t := range times {
        if len(graph[t[0]]) <= 0 {
            graph[t[0]] = make([]edge,0)
        }
        // 题目说有向图，所以不必创建双向edge
        graph[t[0]] = append(graph[t[0]],edge{t[1],t[2]})
    }
    // 距离数字，记录K到i的距离，初始都最大
    dist := make([]int,N+1)
    for i:=0;i<len(dist);i++ {
        dist[i] = math.MaxInt64
    }
    dist[K]=0 // 到本身距离为0
    done := make([]int,N+1) // 记录处理完的节点
    for i:=0;i<N;i++ {
        u := mindis(dist,done)  // 每次从未访问的节点中找距离最小的作为u，
        done[u]=1 // 被访问过
        for _,v := range graph[u] {
            // k到v的距离，是否能被k-u-v替代,找最短路径
            dist[v.v] = min(dist[v.v],dist[u]+v.w)
        }
    }
    ans := 0
    for i:=1;i<len(dist);i++ { //节点从1-N
        ans = max(dist[i],ans)
    }
    if ans == math.MaxInt64 {return -1}
    return ans
}
func mindis(dist,done []int) int {
    mi,mii := math.MaxInt64,0
    for i:=1;i<len(dist);i++ {
        if done[i] <=0 {
            if dist[i]<mi {
                mi,mii = dist[i],i
            }
        }
    }
    return mii
}
func min(a,b int) int {
    if a<b {return a}
    return b
}
func max(a,b int) int {
    if a<b {return b}
    return a
}