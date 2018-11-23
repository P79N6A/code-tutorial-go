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
已有序列都+1并不影响ID关系
*/
func diStringMatch(s string) []int {
    if s=="I" {return []int{0,1}}
    if s=="D" {return []int{1,0}}
    ans := diStringMatch(s[:(len(s)-1)])
    if s[len(s)-1] == 'I' {
        ans = append(ans, len(s))
    } else {
        //如果是D,全都+1还是保持原来顺序,把0 append到最后即可
        for i:=0;i<len(ans);i++ {
            ans[i] += 1
        }// 维持原有的ID关系；都+1是不变的
        ans = append(ans,0)
    }
    return ans
}

func diStringMatch2(S string) []int {
    ans := make([]int,0,len(S)+1)
    left,right := 0,len(S)
    /*
    从0-n的两头填数据，这样填入的不是最大就是最小；新加入的数据对原有的ID结构不产生影响
    */
    for i:=0;i<len(S);i++ {
        if S[i]=='I' {
            ans = append(ans,left)
            left+=1
        } else {
            ans = append(ans,right)
            right-=1
        }
    }
    ans = append(ans,left)
    return ans
}