package main

/*
Given a binary tree with N nodes, each node has a different value from {1, ..., N}.

A node in this binary tree can be flipped by swapping the left child and the right child of that node.

Consider the sequence of N values reported by a preorder traversal starting from the root.  Call such a sequence of N values the voyage of the tree.

(Recall that a preorder traversal of a node means we report the current node's value, then preorder-traverse the left child, then preorder-traverse the right child.)

Our goal is to flip the least number of nodes in the tree so that the voyage of the tree matches the voyage we are given.

If we can do so, then return a list of the values of all nodes flipped.  You may return the answer in any order.

If we cannot do so, then return the list [-1].



Example 1:



Input: root = [1,2], voyage = [2,1]
Output: [-1]
Example 2:



Input: root = [1,2,3], voyage = [1,3,2]
Output: [1]
Example 3:



Input: root = [1,2,3], voyage = [1,2,3]
Output: []


Note:

1 <= N <= 100
*/

func main() {

}

/**
 * Definition for a binary tree node.
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
/*
二叉树，肯定是递归的，比较好拆分子问题
*/
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    ans := make([]int, 0)
    if !dfs(root, voyage, &ans) {
        return []int{-1}
    }
    return ans
}
func dfs(root *TreeNode, vo []int, ans *[]int) bool {
    if root == nil {
        return true
    }
    if len(vo) <= 0 || root.Val != vo[0] {
        *ans = append(*ans, -1)
        return false
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil && root.Right != nil {
        return dfs(root.Right, vo[1:], ans)
    }
    if root.Left != nil && root.Right == nil {
        return dfs(root.Left, vo[1:], ans)
    }
    if root.Left.Val == vo[1] {
        left := 1
        for ; left < len(vo) && vo[left] != root.Right.Val; left++ {}
        return dfs(root.Left, vo[1:left], ans) && dfs(root.Right, vo[left:], ans)
    } else {
        *ans = append(*ans, root.Val)
        left := 1
        for ; left < len(vo) && vo[left] != root.Left.Val; left++ {}
        return dfs(root.Right, vo[1:left], ans) && dfs(root.Left, vo[left:], ans)
    }
}
