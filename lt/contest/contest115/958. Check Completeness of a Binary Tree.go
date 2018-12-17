package main
/*
Given a binary tree, determine if it is a complete binary tree.

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.



Example 1:



Input: [1,2,3,4,5,6]
Output: true
Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
Example 2:



Input: [1,2,3,4,5,null,7]
Output: false
Explanation: The node with value 7 isn't as far left as possible.

Note:

The tree will have between 1 and 100 nodes.
*/
func main() {
    
}
/**
 * Definition for a binary tree node.
*/
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
var heightCache map[*TreeNode]int = make(map[*TreeNode]int)

func isCompleteTree(root *TreeNode) bool {
    return isCompleteTreeLR(root,false)
}
func isCompleteTreeLR(root *TreeNode,force bool) bool {
    /*
    force：是否是满二叉树
    */
    if root == nil {return true}
    lheight := height(root.Left)
    rheight := height(root.Right)
    if force {
        if lheight != rheight {return false}
        return isCompleteTreeLR(root.Left, true) && isCompleteTreeLR(root.Right, true)
    }
    if lheight == rheight {
        //左右一样高，则左边肯定是满的
        return isCompleteTreeLR(root.Left, true) && isCompleteTreeLR(root.Right, false)
    } else if lheight == rheight+1 {
        //左边高，则右边肯定是满的
        return isCompleteTreeLR(root.Left,false) && isCompleteTreeLR(root.Right,true)
    }
    return false
}
func height(r *TreeNode) int {
    if r == nil {return 0}
    if _,ok := heightCache[r];!ok {
        heightCache[r]=max(height(r.Left),height(r.Right))+1
    }
    return heightCache[r]
}
func max(a,b int)int {
    if a>b {return a}
    return b
}