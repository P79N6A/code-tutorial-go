package main

import (
    "fmt"
    "strconv"
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
    paths := make(map[string][]*TreeNode)
    pre(root,paths)
    res := make([]*TreeNode,0)
    for _,v := range paths {
        if len(v) > 1 {
            res = append(res,v[0])
        }
    }
    return res
}
func pre(root *TreeNode,paths map[string][]*TreeNode) string {
    if root == nil {
        return "#"
    }
    p := strconv.Itoa(root.Val)
    lp := pre(root.Left,paths)
    rp := pre(root.Right,paths)
    path := p + lp + rp
    if len(paths[path]) <= 0 {
        paths[path]=make([]*TreeNode,0)
    }
    paths[path]=append(paths[path],root)
    return path
}
func buidtree(idx *int,node []string) *TreeNode{
    if *idx >= len(node) {
        return nil
    }
    if node[*idx] == "#" {
        return nil
    }
    v,_ := strconv.Atoi(node[*idx])
    r := &TreeNode{
        Val:v,
    }
    *idx = *idx+1
    r.Left = buidtree(idx,node)
    *idx = *idx+1
    r.Right = buidtree(idx,node)
    return r
}
func deserializeTree(data []string) *TreeNode {
    idx := 0
    return buidtree(&idx,data)
}
func main() {
    //root := deserializeTree([]string{"0","#","0","#","0","#","0","#","0","#","0","#","0","#","0","#","0"})
    root := deserializeTree([]string{"1","2","4","#","#","#","3","2","4","#","#","#","4","#","#"})
    fmt.Println(findDuplicateSubtrees(root))
}
