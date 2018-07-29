package main

import (
    "fmt"
)

/*
There are n cities connected by m flights. Each fight starts from city u and arrives at v with a price w.

Now given all the cities and fights, together with starting city src and the destination dst, your task is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.

Example 1:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200
Explanation:
The graph looks like this:


The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
Example 2:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
Output: 500
Explanation:
The graph looks like this:


The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
Note:

The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
The size of flights will be in range [0, n * (n - 1) / 2].
The format of each flight will be (src, dst, price).
The price of each flight will be in the range [1, 10000].
k is in the range of [0, n - 1].
There will not be any duplicated flights or self cycles.
*/

func main() {
    fmt.Println(findCheapestPrice(3,[][]int{
        {0,1,100},
        {1,2,100},
        {0,2,500},
    },0,2,1))
    
}
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    // price from src to i, answer is price[dst]
    /*
    price[i] 表示第k次的时候，src->i的最小花费
    题目要求最多K次，则k-1计算完了，再计算k，多一次正好是flights的关联关系
    */
    maxprice := 10001*n // 不能用最大值，否则会溢出
    price := make([]int,n)
    //// 不能使最大值，因为还会变大。 这个很好解决了两个点之间没有关联的情况
    for i:=0;i<len(price);i++ {price[i]=maxprice}
    price[src]=0
    for i:=0;i<=K;i++ {
        nprice := make([]int,n)// K增大一次再继续更新
        // 题目要求最多K次，则下一次使用上一次的结果
        for i:=0;i<len(nprice);i++ {nprice[i]=price[i]}
        for _,f := range flights {
            // 看边是否可以更新
            nprice[f[1]] = min(nprice[f[1]],price[f[0]]+f[2])
        }
        price=nprice
        fmt.Println(price)
    }
    if price[dst]>=maxprice {return -1}
    return price[dst]
}
func min(a,b int) int {
    if a>b{return b}
    return a
}