package main

import (
        "sort"
        "fmt"
)
/*
Given an array A of integers, for each integer A[i] we need to choose either x = -K or x = K, and add x to A[i] (only once).

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.



Example 1:

Input: A = [1], K = 0
Output: 0
Explanation: B = [1]
Example 2:

Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]
Example 3:

Input: A = [1,3,6], K = 3
Output: 3
Explanation: B = [4,6,3]


Note:

1 <= A.length <= 10000
0 <= A[i] <= 10000
0 <= K <= 10000

*/

func main() {
        fmt.Println(smallestRangeII([]int{1,3,6},3))
}
/*
问题,一个数组A,给每个元素 +k,或者-K,问所有操作的可能之后,最大值最小值的差.最小多少
思路:首先肯定是大的-K,小的+K; 排序后,结果就是A[n]-A[0]
假设中间一个结果,i之前所有都+K,i之后的所有都-K,因为所有元素都操作,i属于之前,不属于之后的
则数列变成了  (A[0]+K, A[1]+K, ..., A[i]+K, A[i+1]-K, ..., A[n]-K);想象一下,前边上移,后边下移
左侧部分最大值是A[i]+K,最小值A[0]+K; 右侧部分最大值A[n]-K,最小值A[i+1]-K; 则整个数列的最大值是 max(A[i]+K,A[n]-K)
最小值是min(A[0]+K,A[i+1]-K);更新最终的结果即可.
为什么小点都+K,大的-K, 因为是求最小范围,A[i]+K到A[i+1]-K 肯定是A[i]-K到A[i+1]+K的子集,要往小了弄...
为什么是A[i] A[i+1]这么操作?因为A[i]自己不能又+K,又-K.所以只能相邻元素.
参考 https://leetcode.com/problems/smallest-range-ii/discuss/173389/simple-C++-solution-with-explanation
*/
func smallestRangeII(A []int, K int) int {
        if len(A)<=0 {return 0}
        sort.Ints(A)
        ans := A[len(A)-1]-A[0]
        for i:=0;i<len(A)-1;i++ {
                // 左侧A[0]+K, A[1]+K, ..., A[i]+K
                // 右侧A[i+1]-K, ..., A[n]-K
                imax := max(A[len(A)-1]-K,A[i]+K)// 分割后两个有序序列的最大值
                imin := min(A[0]+K,A[i+1]-K)   //分割后两个有序序列的最小值
                //为什么是min? 因为想要更好的结果.
                ans = min(ans,imax-imin)
        }
        return ans
}
func max(a,b int) int {
        if a>b {return a}
        return b
}
func min(a,b int) int {
        if a<b {return a}
        return b
}
