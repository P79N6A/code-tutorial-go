package main

import "fmt"

/*
An undirected, connected tree with N nodes labelled 0...N-1 and N-1 edges are given.

The ith edge connects nodes edges[i][0] and edges[i][1] together.

Return a list ans, where ans[i] is the sum of the distances between node i and all other nodes.

Example 1:

Input: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
Output: [8,12,6,10,10,10]
Explanation:
Here is a diagram of the given tree:
  0
 / \
1   2
   /|\
  3 4 5
We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
equals 1 + 1 + 2 + 2 + 2 = 8.  Hence, answer[0] = 8, and so on.
Note: 1 <= N <= 10000
*/

func main() {
    fmt.Println(sumOfDistancesInTree(6,[][]int{
        {0,1},{0,2},{2,3},{2,4},{2,5},
    }))
    //fmt.Println(sumOfDistancesInTree1(6,[][]int{
    //    {0,1},{0,2},{2,3},{2,4},{2,5},
    //}))
    
}
/*
这个树有点特殊，是无向图，任意一个节点都可以认为是根。
思路1：任意节点看成根，如果在树里面，
 parent到所有点距离和应该等于所有child的距离和 + 所有child的节点个数。 因为所有的child因为parent多了一层距离而已。
一次后序遍历就可以解决,以各个节点为根分别计算 参看代码sumOfDistanceDfs
*/
func sumOfDistancesInTree(N int, edges [][]int) []int {
    adj := make([][]int, N) // 邻接表
    for _, v := range edges {
        if len(adj[v[0]]) <= 0 {adj[v[0]] = make([]int, 0)}
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    /*
    以任意一点为根，分别计算总的距离
    使用node + parent作为key当缓存，也可以理解，因为这样能够控制树的方向
    */
    ret := make([]int,N)
    cache := make(map[int][]int)
    for i:=0;i<len(ret);i++ {
        ret[i],_ =sumOfDistanceDfs(i,-1,adj,cache)
    }
    return ret
}

func sumOfDistanceDfs(v,parent int,adj [][]int,cache map[int][]int) (int,int) {
    // return sumOfDistance v and sumchild of v
    key := v << 15 + parent
    if _,ok := cache[key];ok {
        return cache[key][0],cache[key][1]
    }
    sumV,cntV := 0,1 // cntV默认就是一个，因为自己
    for _,u := range adj[v] {
        if u != parent { // 用来标识方向，不能往回走，就是通过这个将无向图转换成树的
            s,c := sumOfDistanceDfs(u,v,adj,cache)
            cntV += c
            sumV += s+c // sumV等于子节点的sumOfDistance + 子节点个数
        }
    }
    cache[key] = []int{sumV,cntV}
    return sumV,cntV
}

func sumOfDistancesInTree1(N int, edges [][]int) []int {
    adj := make([][]int,N)
    for _,v := range edges {
        if len(adj[v[0]])<=0 {adj[v[0]]=make([]int,0)}
        adj[v[0]] = append(adj[v[0]],v[1])
        adj[v[1]] = append(adj[v[1]],v[0])
    }
    cnt := make([]int,N) //  初始化cnt
    for i:=0;i<len(cnt);i++ {cnt[i]=1}
    sumdis := make([]int,N)
    postOrder(0,-1,&cnt,&sumdis,adj)
    // 其实只使用sumdis[0]，其他的没有用
    preOrder(0,-1,N,&cnt,&sumdis,adj)
    return sumdis
}
func postOrder(v,parent int, counter,sumdis *[]int,adj [][]int) {
    // 如果是一个有向图，是一个树，有父子之分，那一次递归postOrder就能计算，儿子个数，
    for _,u := range adj[v] {
        if u != parent {
            // 先把parent计算完毕
            postOrder(u,v,counter,sumdis,adj)
            (*counter)[v] += (*counter)[u]
            (*sumdis)[v] += (*counter)[u] + (*sumdis)[u]
        }
    }
}
func preOrder(v,parent,N int, counter,sumdis *[]int,adj [][]int) {
    for _,u := range adj[v] {
        if u != parent {
            // 根据parent重新计算child的。剪掉自己的child，加上剩余的child，因为一次移动产生的变化
            (*sumdis)[u] = (*sumdis)[v] - (*counter)[u] + N - (*counter)[u]
            preOrder(u,v,N,counter,sumdis,adj)
        }
    }
}

