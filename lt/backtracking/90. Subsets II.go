package main

import (
    "fmt"
    "sort"
)

/*

Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

*/
func subsetsWithDup(nums []int) [][]int {
    ret := make([][]int, 0)
    res := make([]int, 0)
    sort.Ints(nums)
    visit := make(map[int]bool)
    comb(nums, 0, visit, &res, &ret)
    return ret
}
func comb(data []int, pos int, visit map[int]bool, res *[]int, ret *[][]int) {
    if pos > len(data) {
        return
    }
    n := make([]int, len(*res))
    copy(n, *res)
    // 所有的可能结果都包含在内了。
    *ret = append(*ret, n)
    for i := pos; i < len(data); i++ {
        if i > 0 && visit[i-1] == false && data[i] == data[i-1] {
            // 如果i-1 没有访问过，并且当前的元素和上一个元素相同，说明不需要处理
            // 等前边的在结果集中在做，否则会出现重复
            continue
        }
        visit[i] = true
        *res = append(*res, data[i])
        comb(data, i+1, visit, res, ret)
        *res = (*res)[:len(*res)-1]
        visit[i] = false
    }
}
func main() {
    fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
