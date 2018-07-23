package main

import (
    "math"
    "fmt"
)

/*
A sequence X_1, X_2, ..., X_n is fibonacci-like if:

n >= 3
X_i + X_{i+1} = X_{i+2} for all i + 2 <= n
Given a strictly increasing array A of positive integers forming a sequence, find the length of the longest fibonacci-like subsequence of A.  If one does not exist, return 0.

(Recall that a subsequence is derived from another sequence A by deleting any number of elements (including none) from A, without changing the order of the remaining elements.  For example, [3, 5, 8] is a subsequence of [3, 4, 5, 6, 7, 8].)



Example 1:

Input: [1,2,3,4,5,6,7,8]
Output: 5
Explanation:
The longest subsequence that is fibonacci-like: [1,2,3,5,8].
Example 2:

Input: [1,3,7,11,12,14,18]
Output: 3
Explanation:
The longest subsequence that is fibonacci-like:
[1,11,12], [3,11,14] or [7,11,18].


Note:

3 <= A.length <= 1000
1 <= A[0] < A[1] < ... < A[A.length - 1] <= 10^9
(The time limit has been reduced by 50% for submissions in Java, C, and C++.)
*/

func main() {
    fmt.Println(lenLongestFibSubseq([]int{1,2,3,4,5,6,7,8}))
    fmt.Println(lenLongestFibSubseq([]int{1,3,7,11,12,14,18}))

}
/*
问题的难点在于需要三个元素才能决定是否组成了fibonacci。想最长递增子序列这种只需要两个元素就能决定，dp的转移方程也较好写出。
如何解决？用dp来保持两个之间的状态。
结果要求三个元素的状态，那dp就维持两个元素状态。计算i,j为fibonacci结尾,那就看如果i,j,k组成了fibonacci，那么 i,j计算方式是
比如dp(i,j)=max(dp(j,k)+1) [i,j,k组成了fibonacci]
或者使用二维dp[i][j]=max(dp[j][k]+1) [i,j,k组成fibonacci]
也可以转换成一维，跟最长递增子序列类似，只不过第三个元素使用hashmap来确认是否有效。
*/
type pair struct {
    i,j int
}
func lenLongestFibSubseq(A []int) int {
    idxMap := make(map[int]int)
    for i:=0;i<len(A);i++ {
        idxMap[A[i]]=i
    }
    ret := 0
    //保持以i,j为结尾的fibonacci最长长度
    cache :=make(map[pair]int)
    for i:=1;i<len(A);i++ {
        for j:=0;j<i;j++ {
            max := -math.MaxInt64
            // 使用j,k 来计算i,j；k的判断使用了hash
            if idx,ok := idxMap[A[i]-A[j]];ok && idx < j {
                // 存在和i,j 组成fibonacci的数字,找最大
                if max < cache[pair{j,idx}] {
                    max = cache[pair{j,idx}]
                }
            }
            if max != -math.MaxInt64 {
                // 有有效结果，max+1,将i加入进来
                cache[pair{i, j}] = max + 1
                if ret < cache[pair{i, j}] {
                    ret = cache[pair{i, j}]
                }
            }
        }
    }
    if ret == 0 {return ret}
    return ret+2 // +2，因为之前的判断都没有计算最初两个值
}
