package main
import (
        "strconv"
        "fmt"
)
/*

We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.

Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

(Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

Example 1:
Input: [1,null,0,0,1]
Output: [1,null,0,null,1]

Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.


Example 2:
Input: [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]



Example 3:
Input: [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]



Note:

The binary tree will have at most 100 nodes.
The value of each node will only be 0 or 1.
*/
type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}
func pruneTree(root *TreeNode) *TreeNode {
        prune(root)
        return root
}
func prune(root *TreeNode) bool {
        // return root, and all zero?
        if root == nil {
                return true
        }
        if root.Left == nil && root.Right == nil {
                return root.Val == 0
        }
        okl := prune(root.Left)
        if okl {
                root.Left = nil
        }
        okr := prune(root.Right)
        if okr {
                root.Right = nil
        }
        return root.Val == 0 && okl && okr

}
func main() {
        root := Deserialization([]string{"1","0","0","#","#","0","#","#","1","0","#","#","1","#","#"})
        data := Serialization(root)
        fmt.Println(data)
        pruneTree(root)
        data = Serialization(root)

        fmt.Println(data)

}

////////////////////
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