package main
/*
Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Example 1:

Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2
Example 2:

Input: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
Follow up:

A solution using O(n) space is pretty straight forward.
Could you devise a constant space solution?
*/
/*
问题，说一个二分查找树有两个数字换了，在不修改树的结构情况下，将这两个数字交换回来。
思路：如果是数组，怎么解决？遍历。两种情况
1.如果只出现了一次 arr[i]>arr[i+1] 则将这两个swap就ok了
2.如果出现了两次arr[i]>arr[i+1],arr[j]>arr[j+1],则swap  arr[i],arr[j+1]即可。
那么tree怎么办？ 中序遍历，并记录preNode即可，这样就能做到和前一个数字比较。
还是那个问题，preNode需要外部变量维护，不能通过形参传递方式维护

*/
func main() {
    
}
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}
var first *TreeNode
var second *TreeNode
var pre *TreeNode  // 全局变量
func dfs(r *TreeNode) {
    if r == nil {return}
    dfs(r.Left) // 中序遍历
    if pre != nil {
        if r.Val < pre.Val { //  出现问题
            if first == nil { // 第一次出现
                first,second = pre,r
            } else {  //  第二次出现
                second = r
                return
            }
        }
    }
    pre = r  // 记录pre，必须是全局变量
    dfs(r.Right)
}
func recoverTree(root *TreeNode)  {
    first,second,pre = nil,nil,nil  // 全局变量多次调用会出错，需要清空
    dfs(root)
    first.Val,second.Val = second.Val,first.Val
}