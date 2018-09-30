package main

import "fmt"

/*
In a given integer array A, we must move every element of A to either list B or list C. (B and C initially start empty.)

Return true if and only if after such a move, it is possible that the average value of B is equal to the average value of C, and B and C are both non-empty.

Example :
Input:
[1,2,3,4,5,6,7,8]
Output: true
Explanation: We can split the array into [1,4,5,8] and [2,3,6,7], and both of them have the average of 4.5.
Note:

The length of A will be in the range [1, 30].
A[i] will be in the range of [0, 10000].
*/

func main() {
    fmt.Println(splitArraySameAverage([]int{1,2,3,4,5,6,7,8}))
}
func splitArraySameAverage(A []int) bool {
    /*
    分成两个子数组a,b  a/k = b/(n-k) = sum/n
    a=k*sum/n 由于A是整数，所以k要满足： k*sum%n == 0
    令a长度较短，因为反正分两份，假设k<=n-k 即可。所以k<=n/2
    */
    n := len(A)
    m := n / 2 // a,b 可以互换，则只处理短的就ok了
    sum := 0
    for _, a := range A {
        sum += a
    }
    cansplit := false
    for i := 1; i <= m; i++ {
        if i*sum%n == 0 { // 先判断是否有这个可能
            cansplit = true
            break
        }
    }
    if !cansplit {
        return false
    }
    /*
    sums[i][j] 表示 A[0-i]数组，任意取j个元素， 加和的所有可能性
    sums[i][j] = sum[i-1][j] U sum[i][j-1]+A[i]
    =>转换到一维
    sum[j] U=sum[j-1]
    */
    sums := make([]map[int]bool, m+1)
    sums[0] = map[int]bool{0: true}
    for i := 0; i < len(A); i++ { // 物品,每次新增一个进来，重新考虑所有weight
        for j := m; j >= 1; j-- { // 价值，01背包，从大到小[空间优化]
            if sums[j] == nil {sums[j] = make(map[int]bool)}
            for k, _ := range sums[j-1] {
                sums[j][k+A[i]]=true
            }
        }
    }
    for k:=1;k<=m;k++ {
        if k*sum%n == 0 { // 有效的K
            if _,ok := sums[k][k*sum/n];ok { // 加和是存在的
                return true
            }
        }
    }
    return false
}
