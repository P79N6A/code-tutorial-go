package main

import (
    "strconv"
    "fmt"
    "math"
)

/*
We have a sorted set of digits D, a non-empty subset of {'1','2','3','4','5','6','7','8','9'}.  (Note that '0' is not included.)

Now, we write numbers using these digits, using each digit as many times as we want.  For example, if D = {'1','3','5'}, we may write numbers such as '13', '551', '1351315'.

Return the number of positive integers that can be written (using the digits of D) that are less than or equal to N.



Example 1:

Input: D = ["1","3","5","7"], N = 100
Output: 20
Explanation:
The 20 numbers that can be written are:
1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
Example 2:

Input: D = ["1","4","9"], N = 1000000000
Output: 29523
Explanation:
We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
In total, this is 29523 integers that can be written using the digits of D.


Note:

D is a subset of digits '1'-'9' in sorted order.
1 <= N <= 10^9
*/

func main() {
    fmt.Println(atMostNGivenDigitSet([]string{"1","4","9"},1000000009))
    fmt.Println(atMostNGivenDigitSet([]string{"7"},91))
    fmt.Println(atMostNGivenDigitSet([]string{"9"},99))
    fmt.Println(atMostNGivenDigitSet([]string{"3","5"},4))
    fmt.Println(atMostNGivenDigitSet([]string{"1","4","7"},6))
}
func atMostNGivenDigitSet11(D []string, N int) int {
    S := strconv.Itoa(N)
    /*
    结果分为几类
    1.位数比N小的，这部分直接 乘法原理 加上去就行了
    2.位数跟N一样的情况
    2.1 第i位上，值跟N相等的情况
    2.2 第i位上，值<N的情况

    dp[i] 表示第i位 非0 情况下，能够跟 string(N)[i]... 之后match上的数字个数
    */
    dp := make([]int,len(S)+1)
    dp[len(S)]=1 // 只影响个位数，并且有值相等情况
    for i:=len(S)-1;i>=0;i-- {
        for j:=0;j<len(D);j++ {
            if D[j] < string(S[i]) {
                // 如果比当前的小，那就 加上 乘法关系
                dp[i] += int(math.Pow(float64(len(D)),float64(len(S)-i)))
            } else if D[j] == string(S[i]) {
                // 如果相等，看之后的位数，dp关系
                dp[i] += dp[i+1]
            } // 如果都比D[j]大，那就是0
        }
    }
    for i:=1;i<len(S);i++ {
        // 把位数 < N的位数情况 的加上去
        dp[0] += int(math.Pow(float64(len(D)),float64(i)))
    }
    return dp[0]
}

func atMostNGivenDigitSet(D []string, N int) int {
    NS := strconv.Itoa(N)
    dp := make([]int,len(NS)+1)
    dp[len(NS)]=1 // 最后一个没用，就是用来加和的
    for i:=len(NS)-1;i>=0;i-- {
        for j:=0;j<len(D);j++ { // D is Sorted
            if D[j] < string(NS[i]) {
                dp[i] += int(math.Pow(float64(len(D)),float64(len(NS)-i-1)))
            } else if D[j] == string(NS[i]) {
                dp[i] += dp[i+1]
            }
        }
    }
    for i:=1;i<len(NS);i++ {
        // 把位数 < N的位数情况 的加上去
        dp[0] += int(math.Pow(float64(len(D)),float64(i)))
    }
    return dp[0]

}
