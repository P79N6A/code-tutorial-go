package main

import "fmt"

/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

func main() {
    fmt.Println(permute([]int{1,2,3}))
    fmt.Println(permute2([]int{1,2,3}))

}
func permute2(nums []int) [][]int {
    result := make([][]int,0)
    pem2(nums,0,&result)
    return result
}

func pem2(nums []int,start int, result *[][]int) {
    if start == len(nums)-1 {
        n := make([]int,len(nums))
        copy(n,nums)
        *result = append(*result,n)
        return
    }
    for i:=start;i<len(nums);i++ {
        /*
        可以这样理解，这个for循环是让所有的元素都在第i个位置待了一次；
        子问题就是除了第一个元素外，剩余的再排列一次，不是swap也行，那就得挪动了,参看pem2的挪动case
        */
        s := nums[start]
        copy(nums[start:],nums[start+1:])
        nums[len(nums)-1]=s
        pem2(nums,start+1,result)
    }
}
func permute(nums []int) [][]int {
    result := make([][]int,0)
    pem(nums,0,&result)
    return result
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
        // 和组合的最大区别，第二个参数传入的是start+1
        pem(nums,start+1,result)
        nums[start],nums[i]=nums[i],nums[start]
    }
}
