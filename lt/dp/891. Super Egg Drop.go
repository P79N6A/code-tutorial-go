package main

import (
    "math"
    "fmt"
)

/*
You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break, and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?



Example 1:

Input: K = 1, N = 2
Output: 2
Explanation:
Drop the egg from floor 1.  If it breaks, we know with certainty that F = 0.
Otherwise, drop the egg from floor 2.  If it breaks, we know with certainty that F = 1.
If it didn't break, then we know with certainty F = 2.
Hence, we needed 2 moves in the worst case to know what F is with certainty.
Example 2:

Input: K = 2, N = 6
Output: 3
Example 3:

Input: K = 3, N = 14
Output: 4


Note:

1 <= K <= 100
1 <= N <= 10000

*/

func main() {
    fmt.Println(superEggDrop(1,2))
    fmt.Println(superEggDrop(2,6))
    fmt.Println(superEggDrop(3,14))
}

/*
K个鸡蛋，N层楼
当处理第x层的时候有两种情况，鸡蛋碎，几点不碎
1，当鸡蛋碎了，则需要考虑x-1层情况，使用k-1个鸡蛋
2，鸡蛋没碎， 那就考虑上边层数N-x，还是用k个鸡蛋[因为M层在上边下边不影响测试最终次数]

dp(K,N)代表最少的测试次数，则dp(K,N) = 1+min{max(dp(K-1,x-1),dp(K,N-x))} 其中 1<=x<=N
其中   dp(K-1,x-1)  表示鸡蛋在第x层碎了；则需要验证x-1层，下边的
          dp(K,N-x) 表示鸡蛋在第x层没碎，需要验证上边的N-x层

计算方法：我们在计算第N层的时候，尝试从每一层x[1<=x<N]把鸡蛋扔下来，
分别看碎和不碎两种情况需要的最少测试次数的最大值【因为要cover最坏的情况，才能说我解决了从第x层扔下的问题】。
然后取最小值，因为我们使用最优的方案，将穷举出来的各个层扔下的最坏情况都得到，取最小值，
也就是需要知道从哪一层扔下来的方法是需要测试次数最少的。
有点极大极小值算法MinMax的意思

优化方法：
solve的复杂度 time O(KN^2)，优化：

优化的关键在 min{max(dp(K-1,x-1),dp(K,N-x))} 这个式子上了。二分查找解决问题。
一个递增一个递减的,而且范围固定，那肯定有交点，是的两个最大值最小，那就是两个值最接近的时候最小了。
上边的交叉图很好的表达了这一点。二分查找能够更快的接近这个结果。 solve2

*/
func superEggDrop(K int, N int) int {
    cache := make([][]int,0)
    for i:=0;i<=K;i++ {
        cache = append(cache,make([]int,N+1))
    }
    return solve2(K,N,&cache)
}
func solve(K,N int,cache *[][]int) int {
    // time O(KN^2)
    if K==1 {return N}
    if N<=1 {return N}
    if (*cache)[K][N] > 0 {
        return (*cache)[K][N]
    }
    min := math.MaxInt64
    for x:=1;x<=N;x++ { // 尝试从各个楼层扔下鸡蛋最坏情况的最小值
        down := solve(K-1,x-1,cache)
        up := solve(K,N-x,cache)
        if min > max(down,up) {min = max(down,up)}
    }
    (*cache)[K][N]=min+1
    return (*cache)[K][N]
}
func max(a,b int)int {
    if a>b {return a}
    return b
}

func solve2(K,N int,cache *[][]int) int {
    /*
    // time O(KNlogN)
    相比solve优化的算法：我们要min(max(down,up))
    这里down随着x的变大变大，up随着x的变小变小。
    所以可以采用二分的方法，才能找到min(max(down,up))最优值。
    */
    if K==1 {return N}
    if N<=1 {return N}
    if (*cache)[K][N] > 0 {
        return (*cache)[K][N]
    }
    min := math.MaxInt64
    l,r := 1,N
    for l<r {
        m := (l+r)/2
        down := solve2(K-1,m-1,cache)
        up := solve2(K,N-m,cache)
        if min > max(down,up) {min = max(down,up)}
        if down == up {
            break
        } else if down < up { // 二分，如果down比较小，说明down,up差不多的点在m的前边
            l = m+1
        } else {
            r = m
        }
    }
    (*cache)[K][N]=min+1
    return (*cache)[K][N]
}