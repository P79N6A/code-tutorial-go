package main

import "fmt"
/*
On an NxN chessboard, a knight starts at the r-th row and c-th column and attempts to make exactly K moves. The rows and columns are 0 indexed, so the top-left square is (0, 0), and the bottom-right square is (N-1, N-1).

A chess knight has 8 possible moves it can make, as illustrated below. Each move is two squares in a cardinal direction, then one square in an orthogonal direction.







Each time the knight is to move, it chooses one of eight possible moves uniformly at random (even if the piece would go off the chessboard) and moves there.

The knight continues moving until it has made exactly K moves or has moved off the chessboard. Return the probability that the knight remains on the board after it has stopped moving.



Example:

Input: 3, 2, 0, 0
Output: 0.0625
Explanation: There are two moves (to (1,2), (2,1)) that will keep the knight on the board.
From each of those positions, there are also two moves that will keep the knight on the board.
The total probability the knight stays on the board is 0.0625.


Note:

N will be between 1 and 25.
K will be between 0 and 100.
The knight always initially starts on the board.
*/

func main() {
    fmt.Println(knightProbability(3,2,0,0))
    fmt.Println(knightProbability(8,30,6,4))
}
func knightProbability(N int, K int, r int, c int) float64 {
    return solve(N,K,r,c,make(map[int]float64))
}
func solve(N int, K int, r,c int,cache map[int]float64) float64 {
    // knight的8种跳跃位置
    direction := [][]int{{-2,-1},{-2,1},{2,-1},{2,1},{-1,-2},{-1,2},{1,-2},{1,2}}
    key := K << 28 + r << 10 + c // K,N范围,一个int可以表示.
    if _,ok := cache[key];ok {return cache[key]}
    if K == 0 {return 1}
    var ans float64
    for _,dir := range direction {
        x,y := r+dir[0],c+dir[1]
        if x>=0&&x<N&&y>=0&&y<N { // 在棋盘范围内
            ans += solve(N,K-1,x,y,cache)
        }
    }
    //如果直接记录int是会超过int64范围限制的;可以直接使用float
    //  float-N = sum(float-(N-1))/N
    cache[key]=ans/8 // 8 possible
    return cache[key]
}