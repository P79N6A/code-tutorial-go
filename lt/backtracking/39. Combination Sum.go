package main

import "fmt"

/*

Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/
func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int,0)
    ret := make([]int,0)
    comb(candidates,0,0,target,&ret,&res)
    return res
}
func comb(can []int, pos int,sum int, target int, ret *[]int, res *[][]int) {
    if sum > target {
        return
    }
    if sum == target {
        n := make([]int,len(*ret))
        copy(n,*ret)
        *res = append(*res,n)
        return
    }
    for i:=pos;i<len(can);i++ {
        *ret = append(*ret,can[i])
        // 可以重复，所以pos参数传入的不是i+1了，是i，唯一的区别
        comb(can,i,sum+can[i],target,ret,res)
        *ret = (*ret)[:len(*ret)-1]
    }
}
func main() {
    fmt.Println(combinationSum([]int{2,3,6,7},7))
}
