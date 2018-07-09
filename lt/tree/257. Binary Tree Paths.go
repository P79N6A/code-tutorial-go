package main

import (
        "strings"
        "strconv"
        "fmt"
)
/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3

*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}
func binaryTreePaths(root *TreeNode) []string {
        path := make([]string,0)
        data := make([]string,0)
        pathdfs(root,&path,&data)
        return data
}
func pathdfs(root *TreeNode,path,data *[]string) {
        if root == nil {
                return
        }
        *path = append(*path,strconv.Itoa(root.Val))
        pathdfs(root.Left,path,data)
        pathdfs(root.Right,path,data)
        if root.Left == nil && root.Right == nil {
                *data = append(*data,strings.Join(*path,"->"))
        }
        *path = (*path)[:len(*path)-1]
}
func main() {
        root := Deserialization([]string{"1","0","0","#","#","0","#","#","1","0","#","#","1","#","#"})
        data := binaryTreePaths(root)
        fmt.Println(data)
}

/////////////////////
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
