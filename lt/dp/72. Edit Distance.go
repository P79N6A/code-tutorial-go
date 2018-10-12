package main

import "fmt"

/*
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

func main() {
    fmt.Println(minDistance("horse","ros"))
    fmt.Println(minDistance("intention","execution"))
    
}
// dp[i][j]=dp[i-1][j-1]                                 if w1[i]==w2[j]
// dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1    if  w1[i]!=w2[j]
//              delete      insert      replace
func minDistance(word1 string, word2 string) int {
    l1,l2 := len(word1),len(word2)
    dp := make([][]int,0)
    for i:=0;i<=l1;i++ { // 方便初始值，都会+1
        dp = append(dp,make([]int,l2+1))
    }
    // 初始的边界条件，否则下边dp过程取最小值会出错
    for i:=1;i<=l1;i++ {dp[i][0]=i} // 也很好理解，这个指的是word2为空，那么word1都是insert的
    for j:=1;j<=l2;j++ {dp[0][j]=j}
    // dp[i][j]=dp[i-1][j-1]                                 if w1[i]==w2[j]
    // dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1    if  w1[i]!=w2[j]
    //              delete      insert      replace
    for i:=1;i<=l1;i++ {
        for j:=1;j<=l2;j++ {
            if word1[i-1]==word2[j-1] {
                dp[i][j]=dp[i-1][j-1]
            } else {
                dp[i][j]=min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1]))+1
            }
        }
    }
    return dp[l1][l2]
}
func min(a,b int) int {
    if a<b {return a}
    return b
}