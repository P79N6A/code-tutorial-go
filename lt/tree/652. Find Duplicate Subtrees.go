package main

import (
    "strconv"
    "strings"
    "fmt"
)

/*
Given a binary tree, return all duplicate subtrees.
For each kind of duplicate subtrees, you only need to return the root node of any one of them.

Two trees are duplicate if they have the same structure with same node values.

Example 1:
        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
The following are two duplicate subtrees:
      2
     /
    4
and
    4
Therefore, you need to return above trees' root in the form of a list.
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    path := make(map[string][]*TreeNode)
    preorder(root,path)
    res := make([]*TreeNode,0)
    for k,v := range path {
        if len(v) > 1 {
            res = append(res, v[0])
        }
        fmt.Println(k,len(v))

    }
    return res
}
func preorder(root *TreeNode,paths map[string][]*TreeNode) string {
    if root == nil {return "#"}
    pl := preorder(root.Left,paths)
    pr := preorder(root.Right,paths)
    path := []string{strconv.Itoa(root.Val),pl,pr}
    key := strings.Join(path,";")
    if _,ok := paths[key];!ok {
        paths[key] = make([]*TreeNode,0)
    }
    paths[key]=append(paths[key],root)
    return key
}
func main() {
    root := deserializeTree([]string{"0","#","0","#","0","#","0","#","0","#","0","#","0","#","0","#","0"})
    //root := deserializeTree([]string{"1","2","4","#","#","#","3","2","4","#","#","#","4","#","#"})
    fmt.Println(findDuplicateSubtrees(root))
}

////////////
func serializeTree(root *TreeNode) []string {
    data := make([]string,0)
    serialize(root,&data)
    return data
}
func serialize(root *TreeNode,data *[]string) {
    if root == nil {
        *data = append(*data,"#")
        return
    }
    *data =append(*data,strconv.Itoa(root.Val))
    serialize(root.Left,data)
    serialize(root.Right,data)
}

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