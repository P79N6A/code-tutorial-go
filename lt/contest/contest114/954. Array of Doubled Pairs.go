package main

import (
    "sort"
    "fmt"
)

/*
Given an array of integers A with even length,
return true if and only if it is possible to reorder it such that A[2 * i + 1] = 2 * A[2 * i] for every 0 <= i < len(A) / 2.



Example 1:

Input: [3,1,3,6]
Output: false
Example 2:

Input: [2,1,2,6]
Output: false
Example 3:

Input: [4,-2,2,-4]
Output: true
Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].
Example 4:

Input: [1,2,4,16,8,4]
Output: false


Note:

0 <= A.length <= 30000
A.length is even
-100000 <= A[i] <= 100000
*/
func main() {
    fmt.Println(canReorderDoubled([]int{3,1,3,6}))
    fmt.Println(canReorderDoubled([]int{2,1,2,6}))
    fmt.Println(canReorderDoubled([]int{4,-2,2,-4}))
    fmt.Println(canReorderDoubled([]int{1,2,4,16,8,4}))
    fmt.Println(canReorderDoubled([]int{0,0,-1,-2,2,4}))
}
/*
corner case还是挺多的，负数处理可以参考非负数。但负数和非负数又不能混在一起。所以直接拆分开。
排序，必须排序，因为要从小到大来，小的如果不存在他的2倍数，直接return false了
hashMap配合使用即可。
*/
func canReorderDoubled(A []int) bool {
    pos := []int{}
    neg := []int{}
    for _, a := range A {
        if a < 0 {
            neg = append(neg, -1*a)
        } else {
            pos = append(pos,a)
        }
    }
    return canReorderDoubledPositive(pos)&&canReorderDoubledPositive(neg)
}
func canReorderDoubledPositive(A []int) bool {
    // 只处理非负数的
    sort.Ints(A)
    dic := make(map[int]int)
    for _, a := range A {
        dic[a] += 1
    }
    for _, a := range A {
        if dic[a] == 0 {
            continue
        }
        dic[a] -= 1
        if dic[2*a] == 0 {
            return false
        }
        dic[2*a] -= 1
    }
    return true
}
