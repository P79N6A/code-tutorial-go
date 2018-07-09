package main

import (
    "fmt"
)

/*
Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.

If there is no non-empty subarray with sum at least K, return -1.



Example 1:

Input: A = [1], K = 1
Output: 1
Example 2:

Input: A = [1,2], K = 4
Output: -1
Example 3:

Input: A = [2,-1,2], K = 3
Output: 3


Note:

1 <= A.length <= 50000
-10 ^ 5 <= A[i] <= 10 ^ 5
1 <= K <= 10 ^ 9

考虑使用单调双向队列。由于sum相减>=k，所以做一个单调递增队列。
每个元素入队一次，入队之前计算跟对头元素相减是否>=K，如果是，队头出队，计算备选；
再看队尾，如果准备入队元素小于队尾，说明队尾已经没用了【队列保留的是最小值，次小。。。】出队；

*/
func shortestSubarray(A []int, K int) int {
    ans := len(A)+1
    sum := make([]int,len(A)+1)
    for i:=0;i<len(A);i++ {sum[i+1]=A[i]+sum[i]}
    queue := make([]int,0)
    fmt.Println(sum)
    for i:=0;i<len(sum);i++ {
        for len(queue) > 0 && sum[i] <= sum[queue[len(queue)-1]]{
            // 保持单调队列特性，先处理队尾。
            queue = queue[:len(queue)-1]
        }
        for len(queue) > 0 && sum[i]-sum[queue[0]] >= K {
            // 队列头可能是备选结果
            if ans > i-queue[0] {
                ans = i-queue[0]
            }
            queue = queue[1:]
        }
        queue = append(queue,i)
    }
    if ans == len(A)+1 {return -1}
    return ans
}
func main() {
    //fmt.Println(shortestSubarray([]int{2,-1,2},3))
    //fmt.Println(shortestSubarray([]int{1,2},4))
    fmt.Println(shortestSubarray([]int{84,-37,32,40,95},167))
}
