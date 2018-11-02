package main
/*
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?

*/

func main() {
    
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    stack := make([]*TreeNode,0) // stack 需要出栈的，维护遍历关系
    ans := make([]int,0)  // 结果，由于stack会出栈，所以不能用。并且ans是从前边插入的。
    stack = append(stack,root)
    for len(stack) > 0 {
        t := stack[len(stack)-1]
        stack = stack[:(len(stack)-1)]
        ans = append([]int{t.Val},ans...) // insert front
        if t.Left != nil {
            stack = append(stack,t.Left)
        }
        if t.Right != nil {
            stack = append(stack,t.Right)
        }
    }
    return ans
}