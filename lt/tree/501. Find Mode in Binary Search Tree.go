package main

import (
        "math"
        "fmt"
)
/*

Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
For example:
Given BST [1,null,2,2],
   1
    \
     2
    /
   2
return [2].

Note: If a tree has more than one mode, you can return them in any order.

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).j
*/
type TreeNode struct {
        Val int
        Left *TreeNode
        Right *TreeNode
}
func findMode(root *TreeNode) []int {
        pre := math.MaxInt64
        cmax :=0
        max := 0
        ret := make([]int,0)
        dfs(root,&pre,&cmax,&max,&ret)
        return ret
}
func dfs(root *TreeNode,pre *int,cmax,max *int,ret *[]int) {
        if root == nil {
                return
        }
        dfs(root.Left,pre,cmax,max,ret)
        if *pre == root.Val {
                *cmax += 1
        } else {
                *cmax = 1
        }
        if *cmax > *max {
                *max = *cmax
                *ret = []int{root.Val}
        } else if *cmax == *max {
                *ret = append(*ret,root.Val)

        }
        *pre = root.Val
        dfs(root.Right,pre,cmax,max,ret)
}

func main() {
        root := &TreeNode{
                Val:1,
                Right:&TreeNode{
                        Val:2,
                        Right:&TreeNode{
                                Val:2,
                        },
                },
        }
        fmt.Println(findMode(root))
}
