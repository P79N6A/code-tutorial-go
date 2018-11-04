package main
/*
Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.

Example 1:
Input:

           1
         /   \
        3     2
       / \     \
      5   3     9

Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
Example 2:
Input:

          1
         /
        3
       / \
      5   3

Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).
Example 3:
Input:

          1
         / \
        3   2
       /
      5

Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
Example 4:
Input:

          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).


Note: Answer will in the range of 32-bit signed integer.
*/

func main() {
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
    depth := depthTree(root) //获取深度,初始化每一层最大下标最小下标数组
    minIdx := make([]int, depth)
    maxIdx := make([]int, depth)
    dfs(root, 0, 1, &minIdx, &maxIdx)
    ans := 1
    for i := 0; i < len(minIdx); i++ {
        if minIdx[i] > 0&&maxIdx[i] > 0 {
            ans = max(ans, maxIdx[i] - minIdx[i]+1)
        }
    }
    return ans
}
func depthTree(root *TreeNode) int {
    if root == nil {return 0}
    return 1 + max(depthTree(root.Left), depthTree(root.Right))
}
func dfs(root *TreeNode, depth, pos int, minIdx, maxIdx *[]int) {
    /*
    depth是深度,pos指的是树[完全二叉树]映射到数组的下标
    pos是根,则做孩子下标是2*pos,右孩子是2*pos+1
    每次看是否可以更新最大下标和最小下标数组
    */
    if root == nil {return}
    if (*minIdx)[depth]<=0 || (*minIdx)[depth]>pos {
        (*minIdx)[depth]=pos // minIdx[depth]没设置或者pos比他小就更新
    }
    if (*maxIdx)[depth]<=0 || (*maxIdx)[depth]<pos {
        (*maxIdx)[depth]=pos
    }
    dfs(root.Left,depth+1,pos*2,minIdx,maxIdx)
    dfs(root.Right,depth+1,pos*2+1,minIdx,maxIdx)
}
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}