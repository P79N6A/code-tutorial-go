package main

import (
    "strconv"
    "fmt"
)
/*
Print a binary tree in an m*n 2D string array following these rules:

The row number m should be equal to the height of the given binary tree.
The column number n should always be an odd number.
The root node's value (in string format) should be put in the exactly middle of the first row it can be put. The column and the row where the root node belongs will separate the rest space into two parts (left-bottom part and right-bottom part). You should print the left subtree in the left-bottom part and print the right subtree in the right-bottom part. The left-bottom part and the right-bottom part should have the same size. Even if one subtree is none while the other is not, you don't need to print anything for the none subtree but still need to leave the space as large as that for the other subtree. However, if two subtrees are none, then you don't need to leave space for both of them.
Each unused space should contain an empty string "".
Print the subtrees following the same rules.
Example 1:
Input:
     1
    /
   2
Output:
[["", "1", ""],
 ["2", "", ""]]
Example 2:
Input:
     1
    / \
   2   3
    \
     4
Output:
[["", "", "", "1", "", "", ""],
 ["", "2", "", "", "", "3", ""],
 ["", "", "4", "", "", "", ""]]
Example 3:
Input:
      1
     / \
    2   5
   /
  3
 /
4
Output:

[["",  "",  "", "",  "", "", "", "1", "",  "",  "",  "",  "", "", ""]
 ["",  "",  "", "2", "", "", "", "",  "",  "",  "",  "5", "", "", ""]
 ["",  "3", "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]
 ["4", "",  "", "",  "", "", "", "",  "",  "",  "",  "",  "", "", ""]]
Note: The height of binary tree is in the range of [1, 10].
*/

func main() {
    //r := &TreeNode{1,&TreeNode{2,nil,&TreeNode{4,nil,nil}},&TreeNode{3,nil,nil}}
    r := &TreeNode{1,&TreeNode{2,nil,nil},nil}
    x := printTree(r)
    for _,xx := range x {
        fmt.Println(xx)
    }
}
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
    if root==nil {return nil}
    lines := depth(root)
    row := rows(lines)
    ans := make([][]string,0)
    for i:=0;i<lines;i++ {
        ans = append(ans,make([]string,row))
    }
    fill(root,0,row,0,&ans)
    return ans
}
func fill(root *TreeNode,start,end int,depth int,ans *[][]string) {
    if root == nil {return}
    x := (start+end)/2
    (*ans)[depth][x]=strconv.Itoa(root.Val)
    fill(root.Left,start,x,depth+1,ans)
    fill(root.Right,x+1,end,depth+1,ans)
}
func rows(n int) int {
    if n== 1 {return 1}
    return rows(n-1)*2+1 // 一共多少列
}
// lens N = 2*lens(N-1)+1
func depth(root *TreeNode) int { //一共多少行
    if root == nil {return 0}
    return max(depth(root.Left),depth(root.Right))+1
}
func max(a,b int) int {
    if a<b {return b}
    return a
}