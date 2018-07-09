package main
/*
We are given a binary tree (with root node root), a target node, and an integer value `K`.

Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.



Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
Output: [7,4,1]
Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.

Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.

Note:

The given tree is non-empty.
Each node in the tree has unique values 0 <= node.val <= 500.
The target node is a node in the tree.
0 <= K <= 1000.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type parent struct {
        r *TreeNode
        rc *TreeNode

}
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
        /*
        先找到所有的parent，在看parent的反方向指定k是否存在，注意parent本身也有可能是答案
        */
        if target == nil {return nil}
        res := make([]parent,0)
        ret := make([]parent,0)
        search(root,target,&ret,&res)
        dis := make([]int,0)
        for i,idx:=len(res)-1,2;i>=0;i,idx=i-1,idx+1 {
                r := res[i]
                if r.r == nil {
                        continue
                }
                if idx-1 == K {
                        dis = append(dis,r.r.Val)
                } else {
                        dfs(r.rc,K-idx,&dis)
                }
        }
        dfs(target,K,&dis)
        return dis
}
func search(root,target *TreeNode, ret *[]parent,res *[]parent) {
        if root == nil {return}
        if root == target {
                n := make([]parent,len(*ret))
                copy(n,*ret)
                *res = n
                return
        }
        *ret = append(*ret,parent{root,root.Right})
        search(root.Left,target,ret,res)
        *ret = (*ret)[:len(*ret)-1]

        *ret = append(*ret,parent{root,root.Left})
        search(root.Right,target,ret,res)
        *ret = (*ret)[:len(*ret)-1]
}

func dfs(root *TreeNode,k int,ret *[]int) {
        if k<0 || root == nil {
                return
        }
        if 0 == k && root != nil {
                *ret = append(*ret,root.Val)
                return
        }
        dfs(root.Left,k-1,ret)
        dfs(root.Right,k-1,ret)
}

func main() {
}
