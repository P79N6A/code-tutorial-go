package main

/*
n个栅栏
k种颜色
不能有三个相同的

如果是不能有四个相同
dp[i]=dp[i-1]*(k-1) + dp[i-2]*(k-1)+dp[i-3]*(k-1)
*/
import "fmt"

func numWays (n int, k int) int {
    if n == 0 {return 0}
    if n == 1 {return k}
    if n == 2 {return k*k}
    // write your code here
    dp := make([]int,n)
    dp[0]=k
    dp[1]=k*k
    for i:=2;i<n;i++ {
        dp[i]=(k-1)*dp[i-1]+dp[i-2]*(k-1)
    }
    return dp[n-1]
}


func main() {
    fmt.Println(numWays(3,2))
    
}
