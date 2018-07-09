package main

import (
    "math"
    "fmt"
)

/*
You have 4 cards each containing a number from 1 to 9.
You need to judge whether they could operated through *, /, +, -, (, ) to get the value of 24.

Example 1:
Input: [4, 1, 8, 7]
Output: True
Explanation: (8-4) * (7-1) = 24
Example 2:
Input: [1, 2, 1, 2]
Output: False
Note:
The division operator / represents real division, not integer division. For example, 4 / (1 - 2/3) = 12.
Every operation done is between two numbers.
In particular, we cannot use - as a unary operator.
For example, with [1, 1, 1, 1] as input, the expression -1 - 1 - 1 - 1 is not allowed.
You cannot concatenate numbers together. For example, if the input is [1, 2, 1, 2],
we cannot write this as 12 + 12.
*/
func judgePoint24(nums []int) bool {
    result := make([][]int,0)
    pem(nums,0,&result)
    for _,r := range result {
        fmt.Println(r)
        if judgePoint24CanNotMoveNums(r) {
            return true
        }
    }
    return false
}
func pem(nums []int,start int, result *[][]int) {
    if start == len(nums)-1 {
        n := make([]int,len(nums))
        copy(n,nums)
        *result = append(*result,n)
        return
    }
    for i:=start;i<len(nums);i++ {
        nums[start],nums[i]=nums[i],nums[start]
        pem(nums,start+1,result)
        nums[start],nums[i]=nums[i],nums[start]
    }
}

func judgePoint24CanNotMoveNums(nums []int) bool {
    ret := dfs(nums,0,4)
    for _,r := range ret {
        if math.Abs(r-float64(24)) < 0.1 {
            return true
        }
    }
    return false
}
func dfs(nums []int, start,end int) []float64 {
    if end-start == 1 {
        return []float64{float64(nums[start])}
    }
    ret := make([]float64,0)
    for i:=start+1;i<end;i++ {
        l := dfs(nums,start,i)
        r := dfs(nums,i,end)
        for _,ll := range l {
            for _,rr := range r {
                ret = append(ret,[]float64{ll-rr,ll+rr,ll*rr,ll/rr}...)
            }
        }

    }
    return ret
}
func main() {
    //fmt.Println(judgePoint24([]int{4,1,8,7}))
    fmt.Println(judgePoint24([]int{8,1,6,6}))

    
}
