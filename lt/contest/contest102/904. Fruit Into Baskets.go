package main
/*
In a row of trees, the i-th tree produces fruit with type tree[i].

You start at any tree of your choice, then repeatedly perform the following steps:

Add one piece of fruit from this tree to your baskets.  If you cannot, stop.
Move to the next tree to the right of the current tree.  If there is no tree to the right, stop.
Note that you do not have any choice after the initial choice of starting tree: you must perform step 1, then step 2, then back to step 1, then step 2, and so on until you stop.

You have two baskets, and each basket can carry any quantity of fruit, but you want each basket to only carry one type of fruit each.

What is the total amount of fruit you can collect with this procedure?



Example 1:

Input: [1,2,1]
Output: 3
Explanation: We can collect [1,2,1].
Example 2:

Input: [0,1,2,2]
Output: 3
Explanation: We can collect [1,2,2].
If we started at the first tree, we would only collect [0, 1].
Example 3:

Input: [1,2,3,2,2]
Output: 4
Explanation: We can collect [2,3,2,2].
If we started at the first tree, we would only collect [1, 2].
Example 4:

Input: [3,3,3,1,2,1,1,2,3,3,4]
Output: 5
Explanation: We can collect [1,2,1,1,2].
If we started at the first tree or the eighth tree, we would only collect 4 fruits.


Note:

1 <= tree.length <= 40000
0 <= tree[i] < tree.length

*/

func main() {
}
//  题目其实是找最多两个不同数字的连续子数组的最大长度,第一次读了半天题目没读懂
func totalFruit(tree []int) int {
        cnt := make(map[int]int)
        ans := 0
        for i,j:=0,0;i<len(tree);i++ { // 左边界每次左移就行
                for j<len(tree) && len(cnt) <= 2 {
                        cnt[tree[j]] += 1
                        j += 1 // 右边界移动
                }
                //超限制了或者结束了
                if len(cnt) <= 2 {
                        ans = max(ans,j-i) // 没超限，结束了
                } else {
                        ans = max(ans,j-i-1) // 超限制了，不算第j个，
                }
                //fmt.Println(cnt,ans,j,i)
                cnt[tree[i]] -= 1  // 左边界左移之前处理
                if cnt[tree[i]] == 0 {
                        delete(cnt,tree[i])
                }
        }
        return ans
}
func max(a,b int) int {
        if a > b {return a}
        return b
}
