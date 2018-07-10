package main

import "fmt"

/*
There is a strange printer with the following two special requirements:

The printer can only print a sequence of the same character each time.
At each turn, the printer can print new characters starting from and ending at any places,
and will cover the original existing characters.
Given a string consists of lower English letters only,
your job is to count the minimum number of turns the printer needed in order to print it.

Example 1:
Input: "aaabbb"
Output: 2
Explanation: Print "aaa" first and then print "bbb".
Example 2:
Input: "aba"
Output: 2
Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.
Hint: Length of the given string will not exceed 100.
*/
//dp[i][j]=dp[i][k] + dp[k+1][j] -1 if char[k]==char[j]
func strangePrinter(s string) int {
    dp := make([][]int,0)
    for i:=0;i<=len(s);i++{
        dp = append(dp,make([]int,len(s)+1))
    }
    for i:=0;i<len(s)-1;i++ {
        for j:=i;j<len(s);j++ {
            for k:=i+1;k<j;k++ {
                if s[k] == s[j] {
                    if dp[i][j] < dp[i][k] + dp[k+1][j]-1 {
                        dp[i][j] = dp[i][k] + dp[k+1][j]-1
                    }
                }
            }
        }
    }
    return dp[0][len(s)]
}
func main() {
    fmt.Println(strangePrinter("aaabbbaccaa"))
}
