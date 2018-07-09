package main
/*
[LeetCode] Binary Tree Longest Consecutive Sequence 二叉树最长连续序列


Given a binary tree, find the length of the longest consecutive sequence path.



The path refers to any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The longest consecutive path need to be from parent to child (cannot be the reverse).

For example,

   1
    \
     3
    / \
   2   4
        \
         5
Longest consecutive sequence path is 3-4-5, so return 3.

   2
    \
     3
    /
   2
  /
 1
Longest consecutive sequence path is 2-3,not3-2-1, so return 2.


*/
func longestConsecutive(root *TreeNode) int {
    if root == nil {return 0}
    max := 0
    dfs(root,root.Val,&max)
    return max
}
func dfs(root *TreeNode,p int, max *int) (int,int) {
    if root == nil {
        return 0,0
    }
    ll,lr := dfs(root.Left,root.Val,max)
    if root.Val == p + 1 {
        lr = lr +1
        ll = 0
    } else if root.Val == p-1 {
        ll = ll + 1
        lr = 0
    } else {
        ll,lr = 0,0
    }

    rl,rr := dfs(root.Right,root.Val,max)
    if root.Val == p + 1 {
        rl = rl +1
        rr = 0
    } else if root.Val == p-1 {
        rr = rr + 1
        rl = 0
    } else {
        rl, rr = 0, 0
    }
    if lr + rr + 1 > *max {
        *max = lr + rr + 1
    }
    if ll + rl > *max {
        *max = ll + rl + 1
    }
    ml := ll
    if rl > ll {
        ml = rl
    }
    mr := lr
    if rr > mr {
        mr = rr
    }
    return ml+1,mr+1
}

func main() {
    
}
