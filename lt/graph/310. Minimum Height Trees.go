package graph

import "fmt"

/*

For a undirected graph with tree characteristics, we can choose any node as the root. The result graph is then a rooted tree.
Among all possible rooted trees, those with minimum height are called minimum height trees (MHTs). Given such a graph, write a function to find all the MHTs and return a list of their root labels.

Format
The graph contains n nodes which are labeled from 0 to n - 1. You will be given the number n and a list of undirected edges (each edge is a pair of labels).

You can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.

Example 1 :

Input: n = 4, edges = [[1, 0], [1, 2], [1, 3]]

        0
        |
        1
       / \
      2   3

Output: [1]
Example 2 :

Input: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]

     0  1  2
      \ | /
        3
        |
        4
        |
        5

Output: [3, 4]
Note:

According to the definition of tree on Wikipedia: “a tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.”
The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
*/
func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {return []int{0}}
    uedges := make([]map[int]bool,n)
    for _,edge := range edges {
        if len(uedges[edge[0]]) == 0 {
            uedges[edge[0]] = make(map[int]bool)
        }
        uedges[edge[0]][edge[1]]=true
        if len(uedges[edge[1]]) == 0 {
            uedges[edge[1]] = make(map[int]bool)
        }
        uedges[edge[1]][edge[0]]=true
    }
    fmt.Println(uedges)

    leaves := make([]int,0)
    for k,v := range uedges {
        if len(v) == 1 {
            leaves = append(leaves,k)
        }
    }
    for {
        newleaves := make([]int,0)
        for _,l := range leaves {
            for v,_ := range uedges[l] {
                delete(uedges[v],l)
                if len(uedges[v]) == 1 {
                    newleaves = append(newleaves,v)
                }
            }
        }
        if len(newleaves) == 0 {
            break
        }
        leaves = newleaves
    }
    return leaves
}
func findMinHeightTreesDfs(n int, edges [][]int) []int {
    uedges := make([][]int,n)
    for _,edge := range edges {
        if len(uedges[edge[0]]) == 0 {
            uedges[edge[0]] = make([]int,0)
        }
        uedges[edge[0]] = append(uedges[edge[0]],edge[1])
        if len(uedges[edge[1]]) == 0 {
            uedges[edge[1]] = make([]int,0)
        }
        uedges[edge[1]] = append(uedges[edge[1]],edge[0])
    }
    min := n+1
    ret := make([]int,0)
    for i:=0;i<n;i++ {
        m := dfsh(uedges,make([]int,n),i)
        if m < min {
            ret = make([]int,0)
            ret = append(ret,i)
            min = m
        } else if m == min {
            ret = append(ret,i)
        }
    }
    return ret
}
func dfsh(edges [][]int,visit []int,i int) int {
    max := 0
    visit[i]=1
    for _,ii := range edges[i] {
        if visit[ii]==1 {
            continue
        }
        n := dfsh(edges,visit,ii)
        if n > max {
            max = n
        }
    }
    return 1+max
}
func main() {
    /*
    fmt.Println(findMinHeightTrees(4,[][]int{
        []int{1,0},
        []int{1,2},
        []int{1,3},
    }))
    */
    // [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]
    /*
    fmt.Println(findMinHeightTrees(6,[][]int{
        []int{0,3},
        []int{1,3},
        []int{2,3},
        []int{4,3},
        []int{5,4},
    }))
    */
    fmt.Println(findMinHeightTrees(1,[][]int{
    }))
}
