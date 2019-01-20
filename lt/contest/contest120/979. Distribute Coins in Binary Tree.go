package main

import (
    "fmt"
    "sort"
)

/*
Given the root of a binary tree with N nodes, each node in the tree has node.val coins, and there are N coins total.

In one move, we may choose two adjacent nodes and move one coin from one node to another.  (The move may be from parent to child, or from child to parent.)

Return the number of moves required to make every node have exactly one coin.



Example 1:



Input: [3,0,0]
Output: 2
Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
Example 2:



Input: [0,3,0]
Output: 3
Explanation: From the left child of the root, we move two coins to the root [taking two moves].  Then, we move one coin from the root of the tree to the right child.
Example 3:



Input: [1,0,2]
Output: 2
Example 4:



Input: [1,0,0,null,3]
Output: 4


Note:

1<= N <= 100
0 <= node.val <= N
*/

func main() {
    t := &TreeNode{
        1,
        &TreeNode{0,nil,nil},
        &TreeNode{2,nil,nil},
    }
    fmt.Println(distributeCoins(t))
}

/**
 * Definition for a binary tree node.
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
func distributeCoins(root *TreeNode) int {
    if root == nil {return 0}
    node := nodeNum(root) // 有多少个节点
    coin := coinNum(root) // 有多少个coin
    ans := abs(coin - node)
    // 使用root作为传输媒体
    // distributeCoins(root.Left) 将左边多出来/或者少出来的挪到root需要都次数
    // distributeCoins(root.Right) 将右边多出来/或者少出来的挪到root需要的次数
    return ans + distributeCoins(root.Left) + distributeCoins(root.Right)
}
func abs(a int) int {
    if a < 0 {
        a = -1*a
    }
    return a
}
func coins(root *TreeNode) int {
    if root == nil {return 0}
    return root.Val-1 + coins(root.Left)+coins(root.Right)
}

func nodeNum(root *TreeNode) int {
    if root == nil {return 0}
    return 1 + nodeNum(root.Left) + nodeNum(root.Right)
}
func coinNum(root *TreeNode) int {
    if root == nil {return 0}
    return root.Val + coinNum(root.Left) + coinNum(root.Right)
}

func distributeCoins1(root *TreeNode) int {
    node := make([]*TreeNode,0)
    parent := make(map[*TreeNode]*TreeNode)
    fillNode(root,nil,&node,parent)
    sort.Slice(node, func(i, j int) bool {
        return node[i].Val < node[j].Val
    })
    ret := 0
    for i:=0;i<len(node);i++ {
        if node[i].Val > 1 {
            visit := make(map[*TreeNode]bool)
            ans := 0
            carry := node[i].Val-1
            node[i].Val = 1
            dfs(node[i],&carry,parent,0,visit,&ans)
            ret += ans
            fmt.Println(node[i].Val,ans)
        }
    }
    return ret
}
func dfs(root *TreeNode,carry *int,parents map[*TreeNode]*TreeNode,depth int,visit map[*TreeNode]bool, ans *int) {
    if *carry == 0 || root == nil{
        return
    }
    if visit[root]==true {
        return
    }
    visit[root]=true
    if root.Val == 0 {
        *carry -= 1
        root.Val = 1
        *ans += depth
    }
    dfs(root.Left,carry,parents,depth+1,visit,ans)
    dfs(root.Right,carry,parents,depth+1,visit,ans)
    dfs(parents[root],carry,parents,depth+1,visit,ans)
}

func fillNode(root,parent *TreeNode,nodes *[]*TreeNode,parents map[*TreeNode]*TreeNode) {
    if root == nil {return}
    *nodes = append(*nodes,root)
    parents[root]=parent
    fillNode(root.Left,root,nodes,parents)
    fillNode(root.Right,root,nodes,parents)
}

