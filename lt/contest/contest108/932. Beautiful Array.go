package main

import "fmt"

/*
For some fixed N, an array A is beautiful if it is a permutation of the integers 1, 2, ..., N, such that:

For every i < j, there is no k with i < k < j such that A[k] * 2 = A[i] + A[j].

Given N, return any beautiful array A.  (It is guaranteed that one exists.)



Example 1:

Input: 4
Output: [2,1,4,3]
Example 2:

Input: 5
Output: [3,1,2,5,4]

1 <= N <= 1000

*/

func main() {
    fmt.Println(beautifulArray(1))
    fmt.Println(beautifulArray(2))
    fmt.Println(beautifulArray(4))
    fmt.Println(beautifulArray(5))
    fmt.Println(beautifulArray(200))
    fmt.Println(beautifulArray1(1000))
}

/*
问题：找到一个N的排列结果，使得：对于任意的k, i<k<j 满足  A[k]*2 != A[i]+A[j]
思路：如果是分治，就看如何拆分子问题。
如果一个奇数+一个偶数，肯定不会等于一个偶数
那么，如果A[1-N]已经是一个结果了。那么
B = A[i]*2
C = A[i]*2-1
组成成的，B+C 是一个长度是2N的满足条件的数组了。
这里面有个case，如果2*x > N了，那去掉就可以了。对最终结果无影响。
下面个问题，比如N是最终的，那么需要拆分的子问题集合是什么规模？
对应2*X,2*X+1两个操作都是什么规模的？
2*X：假设N是偶数， 则x = N/2【比如N={1,2,3,4,5,6},x={1,2,3}即可，这样变成了2，4，6】
    如果N是奇数， x还是=N/2
2*X+1 如果N是技术，x=(N-1)/2【比如N={1,2,3,4,5,6},需要x={1,2,3,4}，这样变成了1,3,5,7,9】
*/
func beautifulArray(N int) []int {
    ans := []int{1}
    for len(ans) < N {
        ret := make([]int,0)
        for _,a := range ans {
            x := 2*a-1
            if x <= N {ret = append(ret,x)}
        }
        for _,a := range ans {
            x := 2*a
            if x <= N {ret = append(ret,x)}
        }
        ans = ret
    }
    return ans
}
func beautifulArray1(N int) []int {
    return solve(N)
}
func solve(N int) []int {
    if N == 1 {
        return []int{1}
    }
    ans := make([]int,N)
    i:=0
    /*
    为什么奇数序列是(n+1)/2。
    比如说本次递归需要N，那他是由多少个x经过 2*x  2*x+1运算得来的？
    也就是(N+1)/2 个奇数，  N/2个偶数
    */
    for _,x := range solve((N+1)/2) {
        ans[i] = 2*x-1 // 前边的奇数序列
        i += 1
    }
    for _,x := range solve(N/2) {
        ans [i] = 2*x  // 偶数序列
        i += 1
    }
    return ans
}