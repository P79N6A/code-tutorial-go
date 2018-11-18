package main

/*
We are given an array A of N lowercase letter strings, all of the same length.

Now, we may choose any set of deletion indices, and for each string, we delete all the characters in those indices.

For example, if we have a string "abcdef" and deletion indices {0, 2, 3}, then the final string after deletion is "bef".

Suppose we chose a set of deletion indices D such that after deletions, each remaining column in A is in sorted order.

Formally, the c-th column is [A[0][c], A[1][c], ..., A[A.length-1][c]]

Return the minimum possible value of D.length.



Example 1:

Input: ["cba","daf","ghi"]
Output: 1
Example 2:

Input: ["a","b"]
Output: 0
Example 3:

Input: ["zyx","wvu","tsr"]
Output: 3


Note:

1 <= A.length <= 100
1 <= A[i].length <= 1000

*/
import "fmt"

func main() {
    fmt.Println(minDeletionSize([]string{"cba","daf","ghi"}))
    fmt.Println(minDeletionSize([]string{"rrjk","furt","guzm"}))
}
//删除的列是固定的....shit. 问题想复杂了
/*
题目说如果想删除列,那么所有字符串该列都删掉.
*/
func minDeletionSize(A []string) int {
    ans := 0
    for j:=0;j<len(A[0]);j++ {
        inorder:=true
        for i:=1;i<len(A);i++ {
            if A[i-1][j] >A[i][j]{inorder=false;break}
        }
        if !inorder {ans += 1}
    }
    return ans
}
