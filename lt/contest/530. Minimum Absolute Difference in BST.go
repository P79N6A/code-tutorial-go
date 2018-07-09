package main

import (
      "math"
      "fmt"
)

/*
Given a binary search tree with non-negative values,
find the minimum absolute difference between values of any two nodes.

Example:

Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
Note: There are at least two nodes in this BST.
*/
/**
 * Definition for a binary tree node.
*/
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}
func getMinimumDifference(root *TreeNode) int {
      if root == nil || (root.Left == nil && root.Right == nil) {return 0}
      min := math.MaxInt64
      last := -1
      dfs(root,&last,&min)
      return min
}
func dfs(node *TreeNode,last *int,min *int) {
      if node == nil {
            return
      }
      dfs(node.Left,last,min)
      if *last == -1 {
            *last = node.Val
      } else {
            if int(math.Abs(float64(node.Val-*last))) < *min {
                  *min = int(math.Abs(float64(node.Val-*last)))
            }
            *last = node.Val
      }
      dfs(node.Right,last,min)
}
func main() {
      root := &TreeNode{
            Val:5,
            Left:&TreeNode{
                  Val:4,
            },
            Right:&TreeNode{
                  Val:7,
            },
      }
      fmt.Println(getMinimumDifference(root))
    
}
