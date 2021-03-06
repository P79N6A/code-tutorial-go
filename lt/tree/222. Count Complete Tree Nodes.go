package main

import "math"

/*
Given a complete binary tree, count the number of nodes.

Note:

Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example:

Input:
    1
   / \
  2   3
 / \  /
4  5 6

Output: 6
*/

func main() {

}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := root, root
    ll, rl := 0, 0
    for l != nil {
        l = l.Left
        ll += 1
    }
    for r != nil {
        r = r.Right
        rl += 1
    }
    if rl == ll {
        // 是否组成了一个满的完全二叉树
        return int(math.Pow(float64(2),float64(rl)))-1
    }
    return 1+countNodes(root.Left)+countNodes(root.Right)
}
