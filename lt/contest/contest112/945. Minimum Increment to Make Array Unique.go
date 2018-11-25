package main

import "fmt"
import "sort"
/*
Given an array of integers A, a move consists of choosing any A[i], and incrementing it by 1.

Return the least number of moves to make every value in A unique.



Example 1:

Input: [1,2,2]
Output: 1
Explanation:  After 1 move, the array could be [1, 2, 3].
Example 2:

Input: [3,2,1,2,1,7]
Output: 6
Explanation:  After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
It can be shown with 5 or less moves that it is impossible for the array to have all unique values.


Note:

0 <= A.length <= 40000
0 <= A[i] < 40000
*/

func main() {
    fmt.Println(minIncrementForUnique([]int{3,2,1,2,1,7}))
}
func minIncrementForUnique(A []int) int {
    if len(A) <= 0 {return 0}
    sort.Ints(A) //排序
    diff := A[0]
    ans:=0
    for i:=0;i<len(A);i++ {
        A[i]-=diff // 排序后,归一化到i,这样就能够跟下标比较了.
        if A[i] < i { // 比下表小,说明要增长到下标位置, ans +=
            ans += i-A[i]
        } else { // 比下标大, 那修改diff,之后所有的数字都需要-=diff
            // 比下标大,相当于上了一个台阶.那就让之后所有的数字都上了这个台阶,跟最终的ans就无关了
            diff += A[i]-i
        }
    }
    return ans
}