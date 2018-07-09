package main

import "fmt"
/*
For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

/**
 * Definition for a binary tree node.
 */
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func buildTree(preorder []int, inorder []int) *TreeNode {
        if len(preorder)<=0 || len(inorder) <= 0 {
                return nil
        }
        root := &TreeNode{
                Val:preorder[0],
        }
        preidx := 0
        leftnode := make(map[int]bool)
        for ;preidx<len(inorder);preidx++ {
                if inorder[preidx] == preorder[0]{
                        break
                }
                leftnode[inorder[preidx]]=true
        }
        inidx := 1
        for ;inidx<len(preorder);inidx++ {
                if leftnode[preorder[inidx]]==false {
                        break
                }
        }
        fmt.Println(inidx,preidx,leftnode)
        root.Left=buildTree(preorder[1:inidx],inorder[:preidx])
        root.Right=buildTree(preorder[inidx:],inorder[preidx+1:])
        return root
}
func main() {
        buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
}
