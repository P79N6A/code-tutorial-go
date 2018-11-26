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
    fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
    fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
    fmt.Println(removeStones([][]int{{0, 0}}))
}

/*
问题的意思是，将所有的行或者列相同的点都放到一组里面。问到底能有多少分组？

解法1：DFS，遍历；看有多少个island。如何表达图关系？
解法2：Union Find，并查集
参考  https://www.cnblogs.com/shadowwalker9/p/5999029.html
*/
func removeStones(stones [][]int) int {
    // 使用stone的下标做uionfindd的key
    dsu := NewDSU(len(stones))
    for i:=0;i<len(stones);i++ {
        for j:=i+1;j<len(stones);j++ {
            if (stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1]){
                dsu.Union(i,j) //两两检查各个点能Union就在一起
            }
        }
    }
    return len(stones)-dsu.Group() // union分成了多少个组
}

type DSU struct {
    parents []int //并查集的索引
}

func NewDSU(size int) *DSU {
    d := &DSU{
        parents: make([]int, size),
    }
    for i := 0; i < len(d.parents); i++ {
        d.parents[i] = i // 初始化DSU makeUnion
    }
    return d
}
func (d *DSU) Group() int {
    // 获取当前并查集有几个集合！
    cnt := 0
    for i:=0;i<len(d.parents);i++ {
        if i == d.parents[i] {
            cnt += 1
        }
    }
    return cnt
}
func (d *DSU) Find(x int) int {
    if x != d.parents[x] {
        // 这个递归过程会优化parent指向
        d.parents[x] = d.Find(d.parents[x])
    }
    return d.parents[x]
}
func (d *DSU) Union(x, y int) {
    //需要考虑x,y的rank，如果能够保证x,y的关系，则不需要了
    x = d.Find(x)
    y = d.Find(y)
    if x == y {
        return
    }
    // x,y 的rank，让小的当parent
    if x < y {
        d.parents[y] = x
    } else {
        d.parents[x] = y
    }
}
