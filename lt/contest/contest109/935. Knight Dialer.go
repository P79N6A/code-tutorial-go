package main
/*
A chess knight can move as indicated in the chess diagram below:

 .



This time, we place our chess knight on any numbered key of a phone pad (indicated above), and the knight makes N-1 hops.  Each hop must be from one key to another numbered key.

Each time it lands on a key (including the initial placement of the knight), it presses the number of that key, pressing N digits total.

How many distinct numbers can you dial in this manner?

Since the answer may be large, output the answer modulo 10^9 + 7.



Example 1:

Input: 1
Output: 10
Example 2:

Input: 2
Output: 20
Example 3:

Input: 3
Output: 46


Note:

1 <= N <= 5000
*/
import "fmt"

func main() {
    fmt.Println(knightDialer(2))
    fmt.Println(knightDialer(3))
}
/*
问题:电话号码盘,每个数字的下一个按钮只能是L形的. 问如果允许按键N次,最多能组成多少不同的数字.
思路:经典的backtracking问题.增加记忆化cache即可.压缩cache key
*/
func knightDialer(N int) int {
    //  i是按键,graph[i]是按键的下一条.
    graph := [][]int{{4,6},{8,6},{7,9},{4,8},{3,9,0},{},{1,7,0},{2,6},{1,3},{2,4}}
    ans := 0
    mod := int(1e9)+7
    cache := make(map[int]int)
    for i:=0;i<10;i++ { // 将0-9所有按键得到结果相加
        ans += solve(i,N-1,graph,cache)
        ans %= mod
    }
    return ans
}
func solve(num,n int,graph [][]int,cache map[int]int) int {
    // num 按键, n第几次, graph 邻接表, cache 记忆化
    if n == 0 {return 1}
    mod := int(1e9)+7
    key := num<<25 + n // num给的从0-9,n最大不超过5k,一个int压缩key正好.
    if _,ok := cache[key];ok {
        return cache[key]
    }
    ans := 0
    for _,v := range graph[num] {
        ans += solve(v,n-1,graph,cache)
        ans %=mod
    }
    cache[key]=ans
    return ans
}