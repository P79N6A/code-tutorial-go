package main

import (
    "math"
    "strconv"
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
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
func getMinimumDifference(root *TreeNode) int {
    pre,res := -1 * math.MaxInt64,math.MaxInt64
    inorder(root,&pre,&res)
    return res
}
func inorder(root *TreeNode,pre,res *int) {
    if root == nil{
        return
    }
    inorder(root.Left,pre,res)
    if *pre != -1 * math.MaxInt64 {
        if root.Val - *pre < *res {
            *res = root.Val-*pre
        }
    }
    fmt.Println(*pre,root.Val)
    *pre = root.Val
    inorder(root.Right,pre,res)
}

func main() {
    root := &TreeNode{
        Val:5,
        Right:&TreeNode{
            Val:7,
        },
        Left:&TreeNode{
            Val:4,
        },
    }
    fmt.Println(getMinimumDifference(root))
}

//////////
func deserializeTree(data []string) *TreeNode {
    idx := 0
    return buildTree(data,&idx)
}
func buildTree(data []string, idx *int) *TreeNode {
    if *idx >= len(data) ||  data[*idx]=="#" {
        return nil
    }
    v,_ := strconv.Atoi(data[*idx])
    node := &TreeNode{
        Val:v,
    }
    *idx += 1
    node.Left = buildTree(data,idx)
    *idx += 1
    node.Right = buildTree(data,idx)
    return node
}