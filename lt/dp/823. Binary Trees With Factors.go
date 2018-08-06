package main

import (
    "sort"
    "fmt"
)

/*
Given an array of unique integers, each integer is strictly greater than 1.

We make a binary tree using these integers and each number may be used for any number of times.

Each non-leaf node's value should be equal to the product of the values of it's children.

How many binary trees can we make?  Return the answer modulo 10 ** 9 + 7.

Example 1:

Input: A = [2, 4]
Output: 3
Explanation: We can make these trees: [2], [4], [4, 2, 2]
Example 2:

Input: A = [2, 4, 5, 10]
Output: 7
Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].


Note:

1 <= A.length <= 1000.
2 <= A[i] <= 10 ^ 9.

*/

func main() {
    fmt.Println(numFactoredBinaryTrees([]int{2,4}))
}
func numFactoredBinaryTrees(A []int) int {
    numMap := make(map[int]int)
    sort.Ints(A)
    for i:=0;i<len(A);i++ {numMap[A[i]]=i}
    dp := make([]int,len(A))
    for i:=0;i<len(dp);i++ {dp[i]=1}
    for i:=1;i<len(A);i++ {
        for j:=0;j<i;j++ {
            if A[i]%A[j] == 0 {
                if _,ok:=numMap[A[i]/A[j]];ok {
                    // 是否存在这样的A[i],A[j],A[i]/A[j]
                    dp[i] += dp[j]*dp[numMap[A[i]/A[j]]]
                }
            }
        }
    }
    sum := 0
    for i:=0;i<len(dp);i++ {sum+=dp[i]}
    return sum
}
