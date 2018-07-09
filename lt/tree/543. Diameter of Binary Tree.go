package main

import "fmt"
/*

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.


*/
type TreeNode struct {
        Val int
        Left *TreeNode
        Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
        max := 0
        dfs(root,&max)
        return max
}
func dfs(root *TreeNode, max *int) int {
        // return max height
        if root == nil {
                return 0
        }
        l := dfs(root.Left,max)
        r := dfs(root.Right,max)
        height := 1
        if l < r {
                height += r
        } else {
                height += l
        }
        if *max < l+r {
                *max = l+r
        }
        return height
}

func main() {
        root := &TreeNode{
                Val:1,
                Left:&TreeNode{
                        Val:2,
                        Left:&TreeNode{
                                Val:4,
                        },
                        Right:&TreeNode{
                                Val:5,
                        },
                },
                Right:&TreeNode{
                        Val:3,
                },
        }
        fmt.Println(diameterOfBinaryTree(root))
}
