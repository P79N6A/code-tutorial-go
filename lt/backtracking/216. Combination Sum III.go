package main

import "fmt"

/*

Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

Note:

All numbers will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Example 2:

Input: k = 3, n = 9
Output: [[1,2,6], [1,3,5], [2,3,4]]
*/
func combinationSum3(k int, n int) [][]int {
    ret := make([][]int,0)
    res := make([]int,0)
    comb([]int{1,2,3,4,5,6,7,8,9},0,0,n,k,&res,&ret)
    return ret
}
func comb(data []int,pos int,sum,target,num int, res *[]int, ret *[][]int) {
    if sum > target {
        return
    }
    if sum == target && len(*res) == num{
        n := make([]int,len(*res))
        copy(n,*res)
        *ret = append(*ret,n)
        return
    }
    for i:=pos;i<len(data);i++ {
        *res = append(*res,data[i])
        comb(data,i+1,sum+data[i],target,num,res,ret)
        *res=(*res)[:len(*res)-1]
    }
}

func main() {
    fmt.Println(combinationSum3(3,9))
    
}
