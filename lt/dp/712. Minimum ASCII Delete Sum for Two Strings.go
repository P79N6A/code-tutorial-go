package main

import "fmt"

/*
Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.

Example 1:
Input: s1 = "sea", s2 = "eat"
Output: 231
Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
Deleting "t" from "eat" adds 116 to the sum.
At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
Example 2:
Input: s1 = "delete", s2 = "leet"
Output: 403
Explanation: Deleting "dee" from "delete" to turn the string into "let",
adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
Note:

0 < s1.length, s2.length <= 1000.
All elements of each string will have an ASCII value in [97, 122].
*/
// dp[i][j]=min(dp[i+1][j],dp[i][j+1])

func minimumDeleteSum(s1 string, s2 string) int {
    dp := make([][]int,0)
    for i:=0;i<=len(s1);i++ {
        dp = append(dp,make([]int,len(s2)+1))
    }
    for i:=len(s1)-1;i>=0;i-- {
        dp[i][len(s2)] = dp[i+1][len(s2)]+int(s1[i])
    }

    for i:=len(s2)-1;i>=0;i-- {
        dp[len(s1)][i] = dp[len(s1)][i+1]+int(s2[i])
    }
    for i:=len(s1)-1;i>=0;i-- {
        for j:=len(s2)-1;j>=0;j-- {
            if s1[i]==s2[j] {
                dp[i][j]=dp[i+1][j+1]
            } else {
                if dp[i+1][j]+int(s1[i]) > dp[i][j+1]+int(s2[j]) {
                    dp[i][j]=dp[i][j+1]+int(s2[j])
                } else {
                    dp[i][j]=dp[i+1][j]+int(s1[i])
                }
            }
        }
    }
    return dp[0][0]
}
func main() {
    fmt.Println(minimumDeleteSum("delete","leet"))
    
}
