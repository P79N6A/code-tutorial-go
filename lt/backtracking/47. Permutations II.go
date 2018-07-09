package main

import "fmt"

/*

Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/
func permuteUnique(nums []int) [][]int {
    ret := make([][]int,0)
    perm(nums,0,&ret)
    return ret
}
func perm(nums []int, pos int, ret *[][]int) {
    if pos >= len(nums) {
        fmt.Println(nums)
        n := make([]int,len(nums))
        copy(n,nums)
        *ret = append(*ret,n)
        return
    }
    ex := make(map[int]bool)
    for i:=pos;i<len(nums);i++ {
        if ex[nums[i]]==false {
            // 如果本次swap有曾经使用过的元素，则不再做这次swap了
            // 因为这样即使swap，也是重复的。交换过来东西相同。
            // 这种之对比nums[i]，不处理pos，因为本层只关注pos位置。
            ex[nums[i]]=true
            nums[i], nums[pos] = nums[pos], nums[i]
            perm(nums, pos+1, ret)
            nums[i], nums[pos] = nums[pos], nums[i]
        }
    }
}

func main() {
    fmt.Println(permuteUnique([]int{2,1,2,1}))

}
