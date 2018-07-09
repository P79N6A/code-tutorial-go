package main

import (
    "strconv"
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
// 9,3,4,#,#,1,#,#,2,#,6,#,#
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
func main() {
    /*
   1
    \
     3
    /
   2
    */
    root := &TreeNode{
        Val:1,
        Right:&TreeNode{
            Val:3,
            Left:&TreeNode{
                Val:2,
            },
        },
    }
    data := serializeTree(root)
    fmt.Println(data)
    
}
