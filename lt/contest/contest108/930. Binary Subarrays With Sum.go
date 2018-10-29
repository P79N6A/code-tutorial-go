package main
/*
In an array A of 0s and 1s, how many non-empty subarrays have sum S?



Example 1:

Input: A = [1,0,1,0,1], S = 2
Output: 4
Explanation:
The 4 subarrays are bolded below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]


Note:

A.length <= 30000
0 <= S <= A.length
A[i] is either 0 or 1.
*/

import "fmt"

func main() {
    fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1},2))
    fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0},0))
}
func numSubarraysWithSum(A []int, S int) int {
    ans := 0
    sd := make(map[int]int) //  sum dict
    sd[0]=1
    sum := 0
    for _,a := range A {
        sum += a
        if _,ok := sd[sum-S];ok {
            ans += sd[sum-S] // 加上曾经出现过的sum-S的次数
        }
        sd[sum]+=1
    }
    return ans
}