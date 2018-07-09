package main

import "fmt"

/*

Initially on a notepad only one character 'A' is present.
You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted.
Output the minimum number of steps to get n 'A'.

Example 1:
Input: 3
Output: 3
Explanation:
Intitally, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.
Note:
The n will be in the range [1, 1000].
*/
/*
dp[i]=min(dp[i/j]+j)
*/
func minSteps(n int) int {
    if n <= 1 {
        return 0
    }
    dp := make([]int,n+1)
    for i:=0;i<=n;i++ {
        dp[i]=i
        for j:=i/2;j>0;j-- {
            if i%j == 0 {
                if dp[i]>dp[j]+i/j {
                    dp[i]=dp[j] + i/j
                }
            }
        }
    }
    fmt.Println(dp)
    return dp[n]
}
func main() {
    fmt.Println(minSteps(3))

    
}
