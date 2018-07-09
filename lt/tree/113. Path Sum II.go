package main
/*
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1
Return:

[
   [5,4,11,2],
   [5,8,4,5]
]

*/
type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}


func pathSum(root *TreeNode, sum int) [][]int {
        path := make([]int,0)
        data := make([][]int,0)
        sumdfs(root,&path,&data,sum)
        return data
}
func sumdfs(root *TreeNode,path *[]int,data *[][]int,sum int) {
        if root == nil {
                return
        }
        *path = append(*path,root.Val)
        sumdfs(root.Left,path,data,sum)
        sumdfs(root.Right,path,data,sum)
        if root.Left == nil && root.Right == nil {
                s := 0
                for _,p := range *path {
                        s += p
                }
                if s == sum {
                        np := make([]int,len(*path))
                        copy(np,*path)
                        *data = append(*data,np)
                }
        }
        *path = (*path)[:len(*path)-1]
}

func main() {
        pathSum(nil,1)
}
