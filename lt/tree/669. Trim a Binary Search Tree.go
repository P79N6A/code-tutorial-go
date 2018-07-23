package main

import (
    "fmt"
    "strconv"
)

func main() {
    root := &TreeNode{
        Val:1,
        Left:&TreeNode{Val:0},
        Right:&TreeNode{Val:2},
    }
    r := trimBST(root,1,2)
    fmt.Println(Serialization(r))
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root==nil{return nil}
    if root.Val < L {
        return trimBST(root.Right,L,R)
    }
    if root.Val > R {
        return trimBST(root.Left,L,R)
    }
    root.Left = trimBST(root.Left,L,R)
    root.Right = trimBST(root.Right,L,R)
    return root
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func Serialization(root *TreeNode) []string {
    data := make([]string,0)
    dfs(root,&data)
    return data
}
func dfs(root *TreeNode,data *[]string) {
    if root == nil {
        *data = append(*data,"#")
        return
    }
    *data = append(*data,strconv.Itoa(root.Val))
    dfs(root.Left,data)
    dfs(root.Right,data)
}