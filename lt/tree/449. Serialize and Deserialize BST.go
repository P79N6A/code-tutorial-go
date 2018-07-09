package main

import (
        "fmt"
        "strconv"
)
/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary search tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary search tree can be serialized to a string and this string can be deserialized to the original tree structure.

The encoded string should be as compact as possible.

Note: Do not use class member/global/static variables to store states. Your serialize and deserialize algorithms should be stateless.


*/
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
func Deserialization(data []string) *TreeNode {
        idx := 0
        return buildTree(data,&idx)
}
func buildTree(data []string,idx *int) *TreeNode {
        if *idx > len(data) || data[*idx] == "#" {
                return nil
        }
        d,_ := strconv.Atoi(data[*idx])
        node := &TreeNode{
                Val:d,
        }
        *idx += 1
        node.Left = buildTree(data,idx)
        *idx += 1
        node.Right = buildTree(data,idx)
        return node
}
func main() {
        /*
   1
 /   \
2     3
 \
  5
        */
        root := &TreeNode{
                Val:1,
                Left:&TreeNode{
                        Val:2,
                        Right:&TreeNode{
                                Val:5,
                        },
                },
                Right:&TreeNode{
                        Val:3,
                },
        }
        data := Serialization(root)
        fmt.Println(data)
        aroot := Deserialization(data)
        //root.Left.Right.Val=2
        fmt.Println(diff(root,aroot))
}
func diff(root1,root2 *TreeNode) bool {
        if root1 == nil && root2 == nil {
                return true
        }
        if root1 != nil && root2 != nil && root1.Val != root2.Val {
                return false
        }
        left := diff(root1.Left,root2.Left)
        if left == false {
                return false
        }
        right := diff(root1.Right,root2.Right)
        if right == false {return false}
        return true

)
