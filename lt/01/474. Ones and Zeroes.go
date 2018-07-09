package main

import (
    "fmt"
)

/*
In the computer world, use restricted resource you have to generate maximum benefit is
what we always want to pursue.

For now, suppose you are a dominator of m 0s and n 1s respectively.
On the other hand, there is an array with strings consisting of only 0s and 1s.

Now your task is to find the maximum number of strings that you can form with given m 0s and n 1s.
Each 0 and 1 can be used at most once.

Note:
The given numbers of 0s and 1s will both not exceed 100
The size of given string array won't exceed 600.
Example 1:
Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
Output: 4

Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s,
which are “10,”0001”,”1”,”0”
Example 2:
Input: Array = {"10", "0", "1"}, m = 1, n = 1
Output: 2

Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".

*/
// dp[i][j]= i's 0 j's max element
// dp[i][j]=max(dp[i][j],dp[i-x0][j-x1]+1)
//

func findMaxForm(strs []string, m int, n int) int {
    //return dfs(strs,0,m,n)
    dp := make([][]int,0)
    for i:=0;i<=m;i++ {
        dp = append(dp,make([]int,n+1))
    }
    for _,s := range strs {
        for i:=m;i>=0;i-- {
            for j:=n;j>=0;j-- {
                z,o := 0,0
                for i:=0;i<len(s);i++ {
                    if s[i]=='0' {z+=1}
                    if s[i]=='1' {o+=1}
                }
                if i>=z&&j>=o {
                    if dp[i][j] < dp[i-z][j-o]+1 {
                        dp[i][j] = dp[i-z][j-o] + 1
                    }
                }
            }
        }
    }
    return dp[m][n]
}
func dfs(strs []string,start,m,n int) int {
    if m < 0 || n < 0 {
        return -1
    }
    if m >= 0 && n >= 0 && start == len(strs){
        return 0
    }
    if start >= len(strs) {
        return 0
    }
    o,z := 0,0
    for i:=0;i<len(strs[start]);i++ {
        if strs[start][i]=='0' {z+=1}
        if strs[start][i]=='1' {o+=1}
    }
    a :=  dfs(strs,start+1,m,n)
    b := dfs(strs,start+1,m-z,n-o)
    if  b + 1 > a {
        return b+1
    }
    return a
}
func main() {
    fmt.Println(findMaxForm([]string{"1", "0"},1,1)) // 2
    fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"},5,3)) // 4
    fmt.Println(findMaxForm([]string{"10","0001","111001","1","0"},3,4))  // 3
    fmt.Println(findMaxForm([]string{"00","000"},1,10))
    fmt.Println(findMaxForm([]string{"0","0","1","1"},2,2))

}
