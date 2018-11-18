package main
/*
Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

If S[i] == "I", then A[i] < A[i+1]
If S[i] == "D", then A[i] > A[i+1]


Example 1:

Input: "IDID"
Output: [0,4,1,3,2]
Example 2:

Input: "III"
Output: [0,1,2,3]
Example 3:

Input: "DDI"
Output: [3,2,0,1]


Note:

1 <= S.length <= 10000
S only contains characters "I" or "D".
*/

import "fmt"

func main() {
    fmt.Println(diStringMatch("IDID"))
}
/*
如何保证原有的DI关系不变?这样才能线性递归下去...
好像有点排列的意思
*/
func diStringMatch(S string) []int {
    return solve(S)
}
func solve(s string) []int {
    if s=="I" {return []int{0,1}}
    if s=="D" {return []int{1,0}}
    ans := solve(s[:(len(s)-1)])
    if s[len(s)-1] == 'I' {
        ans = append(ans, len(s))
    } else {
        //如果是D,全都+1还是保持原来顺序,把0 append到最后即可
        for i:=0;i<len(ans);i++ {
            ans[i] += 1
        }
        ans = append(ans,0)
    }
    return ans
}