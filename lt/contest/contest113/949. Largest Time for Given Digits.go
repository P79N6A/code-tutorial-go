package main
/*
Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5.  If no valid time can be made, return an empty string.



Example 1:

Input: [1,2,3,4]
Output: "23:41"
Example 2:

Input: [5,5,5,5]
Output: ""


Note:

A.length == 4
0 <= A[i] <= 9
*/

import "fmt"

func main() {
    fmt.Println(largestTimeFromDigits([]int{3,3,4,5}))
}
/*
排列4个数字所有的可能;看有效的

可以四层否循环搞定。或者使用排列算法
*/
func largestTimeFromDigits(A []int) string {
    /*
    19:59
    23:59
    0,1,2,3
    */
    ans := ""
    for _,seq := range permute2(A) {
        if seq[0]*10 + seq[1] < 24 && seq[2]*10+seq[3] < 60 {
            x := fmt.Sprintf("%d%d:%d%d",seq[0],seq[1],seq[2],seq[3])
            if ans == "" || x > ans {
                ans = x
            }
        }
    }
    return ans
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