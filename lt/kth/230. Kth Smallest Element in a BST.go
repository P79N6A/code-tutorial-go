package main
/*
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to
find the kth smallest frequently? How would you optimize the kthSmallest routine?
*/
/**
 * Definition for a binary tree nodeo.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    node := root
    for node != nil {
        c := count(node.Left)
        if c == k-1 {
            return node.Val
        } else if c >= k {
            node = node.Left
        } else {
            node = node.Right
            k = k - c - 1
        }
    }
    return -1
}
/*
递归计算有多少个孩子节点
*/
func count(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + count(root.Left) + count(root.Right)
}


func main() {
    
}
