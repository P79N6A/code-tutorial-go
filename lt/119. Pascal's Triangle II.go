package main

import "fmt"
/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
*/

func getRow(rowIndex int) []int {
        if rowIndex == 0 {return []int{1}}
        if rowIndex == 1 {return []int{1,1}}
        res := make([]int,rowIndex+1)
        res[0]=1
        res[1]=1
        for i := 2;i <= rowIndex;i++ {
                for j := i;j > 0;j-- {
                        res[j] = res[j-1]+res[j]
                }
        }
        return res
}

func main() {
        fmt.Println(getRow(5))
}
