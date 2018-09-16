package main

import "fmt"
/*
Given an array of integers A, find the sum of min(B), where B ranges over every (contiguous) subarray of A.

Since the answer may be large, return the answer modulo 10^9 + 7.



Example 1:

Input: [3,1,2,4]
Output: 17
Explanation: Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.  Sum is 17.


Note:

1 <= A <= 30000
1 <= A[i] <= 30000
*/

func main() {
        fmt.Println(sumSubarrayMins1([]int{1,4,2,1,3}))
}
func sumSubarrayMins1(A []int) int {
        //使用单调栈之后大约有10倍性能提升,O(N)
        //问题是在解题过程中能够想到单调栈,找左边/右边第一个大于/小于给定target的元素
        if len(A) <= 0 {return 0}
        incStack := make([]int,0) // 单调递增栈
        leftSide := make([]int,len(A))
        for i:=0;i<len(A);i++ {
                for len(incStack) > 0 && A[incStack[len(incStack)-1]] > A[i] {
                        // 出栈逻辑,leftside只考虑>A[i]
                        incStack = incStack[:(len(incStack)-1)]
                }
                if len(incStack) > 0 { //在入栈的时候计算leftside
                        leftSide[i] = i - incStack[len(incStack) - 1]-1
                } else {
                        leftSide[i]=i
                }
                incStack = append(incStack,i) // 入栈
        }
        incStack = make([]int,0)
        rightSide := make([]int,len(A))
        for i:=len(A)-1;i>=0;i-- {
                for len(incStack) > 0 && A[incStack[len(incStack)-1]] >= A[i] {
                        // 出栈逻辑,leftside考虑>A[i]
                        incStack = incStack[:(len(incStack)-1)]
                }
                if len(incStack) > 0 { //在入栈的时候计算leftside
                        rightSide[i] = incStack[len(incStack) - 1]-i-1
                } else {
                        rightSide[i]=len(A)-1-i
                }
                incStack = append(incStack,i)
        }
        ret := 0
        mod := int(1e9) + 7
        for i:=0;i<len(A);i++ {
                ln,rn := leftSide[i],rightSide[i]
                ret += A[i] + A[i]*ln%mod + A[i]*rn%mod + ((A[i]*ln)%mod)*rn%mod
                ret %=mod
        }
        return ret%mod
}
func sumSubarrayMins(A []int) int {
        /*
        思路： ret = sum(A[i] * f[i])
        f[i]表示，以A[i]为最小值的子序列个数。怎么计算？
        A[i]本身
        以A[i]为左边的个数 right
        以A[i]为右边的个数 left
        A[i]不是边界的时候， right*left
        考虑有重复，则计算左边界不算等于，右边界计算相等情况
        之后这个问题就变成了,针对A[i]如何找到左边第一个<=A[i]的,右边第一个<A[i]的 => 单调栈
        */
        if len(A) <= 0 {return 0}
        ret := 0
        mod := int(1e9) + 7
        for i:=0;i<len(A);i++ {
                ln,rn := 0,0
                for l:=i-1;l>=0&&A[i]<A[l];l-- { //左开
                        ln += 1
                }
                for r:=i+1;r < len(A)&&A[i]<=A[r];r++ {  // 右闭有等号，否则如果有重复元素会重复计算子序列区间
                        rn += 1
                }
                // 自身,左侧,右侧,中间
                ret += A[i] + A[i]*ln%mod + A[i]*rn%mod + ((A[i]*ln)%mod)*rn%mod
                ret %=mod
        }
        return ret%mod
}

