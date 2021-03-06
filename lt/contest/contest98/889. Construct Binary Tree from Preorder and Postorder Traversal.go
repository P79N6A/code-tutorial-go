package main

import "fmt"

func main() {
        fmt.Println(constructFromPrePost([]int{1,2,4,5,3,6,7},[]int{4,5,2,6,7,3,1}))
}

// Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func constructFromPrePost(pre,post []int) *TreeNode {
    if len(pre)<=0 {return nil}
    if len(pre) == 1 {return &TreeNode{pre[0],nil,nil}}
    root := &TreeNode{pre[0],nil,nil}
    i := 0
    // 通过前序第一个在后序中位置就能判断左子树序列
    for pre[1]!=post[i]{i++}
    root.Left = constructFromPrePost(pre[1:i+2],post[:i+1])
    root.Right = constructFromPrePost(pre[i+2:],post[i+1:len(post)-1])
    return root
}
func solve(pre,post []int) *TreeNode {
        if len(pre)<=0 {return nil}
        if len(pre) == 1 {return &TreeNode{pre[0],nil,nil}}
        root := &TreeNode{pre[0],nil,nil}
        pred := make([]int,32)
        postd := make([]int,32)
        prei,posti := 1,0
        for {
                pred[pre[prei]],postd[post[posti]]=1,1
                prei += 1
                posti += 1
                if same(pred,postd) {
                        break
                }
        }
        root.Left = solve(pre[1:prei],post[:posti])
        root.Right = solve(pre[prei:],post[posti:len(post)-1])
        return root
}
func same(a,b []int) bool {
        for i:=0;i<len(a);i++ {
                if a[i] != b[i] {return false}
        }
        return true
}
