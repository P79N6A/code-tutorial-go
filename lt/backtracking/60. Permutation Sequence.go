package main

import "fmt"

/*
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
*/

func main() {
    fmt.Println(getPermutation(3,2)) // 132
    //fmt.Println(getPermutation(3,3))
    //fmt.Println(getPermutation(4,9))
    //fmt.Println(getPermutation(4,10))
}
func getPermutation(n int, k int) string {
    factorial := make([]int, n+1)
    factorial[0] = 1
    nums := make([]int, n)
    for i := 1; i <= n; i++ {
        factorial[i] = factorial[i-1] * i
        nums[i-1] = i
    }
    // nums = {1,2,3,4,5}
    // 下标是从0开始的，但数字是从1开始，以为k/factorial==0,就把第一个给出去
    k -= 1  // why? 0-base 方便取余，除法
    ret := ""
    /*
    看是否大于(n-1)! 看大了多少个？ 有多少个就算未使用数组第几个数字
    修正k，减去对应次数
    每次都看最大的，如果不够大则肯定是第一个了
    */
    for i := 1; i <= n; i++ {
        // 看n-i的阶乘算不算，从最大的开始，如果不超过，看肯定是第一个idx=0给出去
        idx := k / factorial[n-i]
        ret += string('0' + nums[idx])
        nums = append(nums[:idx], nums[idx+1:]...) //remove the number
        // 下边两种写法等价
        //k -= idx*factorial[n-i] // 去掉减去部分
        k = k%factorial[n-i] // 去掉减去部分
    }
    return ret
}
