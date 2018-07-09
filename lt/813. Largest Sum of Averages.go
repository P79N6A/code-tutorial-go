/*
We partition a row of numbers A into at most K adjacent (non-empty) groups, then our score is the sum of the average of each group. What is the largest score we can achieve?

Note that our partition must use every number in A, and that scores are not necessarily integers.

Example:
Input:
A = [9,1,2,3,9]
K = 3
Output: 20
Explanation:
The best choice is to partition A into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
We could have also partitioned A into [9, 1], [2], [3, 9], for example.
That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.


Note:

1 <= A.length <= 100.
1 <= A[i] <= 10000.
1 <= K <= A.length.
Answers within 10^-6 of the correct answer will be accepted as correct.
*/
package main

import "fmt"

func largestSumOfAverages1(A []int, K int) float64 {
        // dp[i][k] = max(dp[j][k-1]+avg(sum[j-i])) 0<j<i
        alen := len(A)
        dp := make([][]float64, alen+1)
        for i := 0; i <= alen; i++ {
                dp[i] = make([]float64, K+1)
        }
        sumA := make([]int, alen+1)
        for i := 1; i <= alen; i++ {
                sumA[i] += sumA[i-1] + A[i-1]
                dp[i][1]= float64(sumA[i])/float64(i)
        }

        for k := 2; k <= K; k++ {
                for i := k; i <= alen; i++ {
                        for j := k-1; j < i; j++ {
                                avg := float64(sumA[i] - sumA[j]) / float64(i - j)
                                fmt.Printf("k %d,i:%d,j:%d,avg:%f;dp[%d][%d]=%f\n",k,i,j,avg,j,k-1,dp[j][k-1])
                                if dp[i][k] < dp[j][k - 1] + avg {
                                        dp[i][k] = dp[j][k - 1] + avg
                                }
                        }
                }
        }
        //*/

        fmt.Println("xxxxxxx")
        for _, s := range dp {
                fmt.Println(s)
        }
        return dp[alen][K]
}
func largestSumOfAverages(A []int, K int) float64 {
        // dp[i][k] = max(dp[j][k-1]+avg(sum[j-i])) 0<j<i
        alen := len(A)
        if alen <= 0 {
                return 0
        }
        sumA := make([]int, alen+1)
        for i := 0; i < alen; i++ {
                sumA[i+1] += sumA[i] + A[i]
        }
        fmt.Println(sumA)
        dp := make([][]float64, alen+1)
        for i := 0; i <= alen; i++ {
                dp[i] = make([]float64, K+1)
                dp[i][1]= float64(sumA[i])/float64(i)
        }

        for k := 2; k <= K; k++ {
                for i := k; i <= alen; i++ {
                        for j := k-1; j < i; j++ {
                                avg := float64(sumA[i] - sumA[j]) / float64(i - j)
                                fmt.Printf("k %d,i:%d,j:%d,avg:%f;dp[%d][%d]=%f\n",k,i,j,avg,j,k-1,dp[j][k-1])
                                if dp[i][k] < dp[j][k - 1] + avg {
                                        dp[i][k] = dp[j][k - 1] + avg
                                }
                        }
                }
        }
        //*/

        fmt.Println("xxxxxxx")
        for _, s := range dp {
                fmt.Println(s)
        }
        return dp[alen][K]
}
func main() {
        // 20.5
        fmt.Println(largestSumOfAverages([]int{1,2,3,4,5,6,7},4))
        //fmt.Println(largestSumOfAverages([]int{1, 2}, 2))
        //fmt.Println(largestSumOfAverages([]int{9,1,2,3,9},3))
}
