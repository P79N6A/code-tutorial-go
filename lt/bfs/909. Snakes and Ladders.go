package main
/*
On an N x N board, the numbers from 1 to N*N are written boustrophedonically starting from the bottom left of the board, and alternating direction each row.  For example, for a 6 x 6 board, the numbers are written as follows:

36 35 34 33 32 31
25 26 27 28 29 30
24 23 22 21 20 19
13 14 15 16 17 18
12 11 10 09 08 07
01 02 03 04 05 06
You start on square 1 of the board (which is always in the last row and first column).  Each move, starting from square x, consists of the following:

You choose a destination square S with number x+1, x+2, x+3, x+4, x+5, or x+6, provided this number is <= N*N.
If S has a snake or ladder, you move to the destination of that snake or ladder.  Otherwise, you move to S.
A board square on row r and column c has a "snake or ladder" if board[r][c] != -1.  The destination of that snake or ladder is board[r][c].

Note that you only take a snake or ladder at most once per move: if the destination to a snake or ladder is the start of another snake or ladder, you do not continue moving.

Return the least number of moves required to reach square N*N.  If it is not possible, return -1.

Example 1:

Input: [
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,-1,-1,-1,-1,-1],
[-1,35,-1,-1,13,-1],
[-1,-1,-1,-1,-1,-1],
[-1,15,-1,-1,-1,-1]]
Output: 4
Explanation:
At the beginning, you start at square 1 [at row 5, column 0].
You decide to move to square 2, and must take the ladder to square 15.
You then decide to move to square 17 (row 3, column 5), and must take the snake to square 13.
You then decide to move to square 14, and must take the ladder to square 35.
You then decide to move to square 36, ending the game.
It can be shown that you need at least 4 moves to reach the N*N-th square, so the answer is 4.
Note:

2 <= board.length = board[0].length <= 20
board[i][j] is between 1 and N*N or is equal to -1.
The board square with number 1 has no snake or ladder.
The board square with number N*N has no snake or ladder.

*/

import (
        "fmt"
        "math"
)
/*

36 35 34 33 32 31
25 26 27 28 29 30
24 23 22 21 20 19
13 14 15 16 17 18
12 11 10 09 08 07
01 02 03 04 05 06
*/
func main() {
        fmt.Println(snakesAndLadders([][]int{
                {-1,-1,-1,-1,-1,-1},
                {-1,-1,-1,-1,-1,-1},
                {1,-1,-1,-1,-1,-1},
                {-1,35,-1,-1,13,-1},
                {-1,-1,-1,-1,-1,-1},
                {-1,15,-1,-1,-1,-1}}))
        fmt.Println(snakesAndLadders([][]int{
                {1,1,-1},
                {1,1,1},
                {-1,1,1}}))
        fmt.Println(snakesAndLadders([][]int{
                {-1,-1,-1},
                {-1,9,8},
                {-1,8,9}}))
        fmt.Println(snakesAndLadders([][]int{
                {-1,1,2,-1},
                {2,13,15,-1},
                {-1,10,-1,-1},
                {-1,6,2,8}})) // 2 //*/
/*
21 22 23 24 25
20 19 18 17 16
11 12 13 14 15
10 9 8 7 6
1 2 3 4 5
*/
        fmt.Println(snakesAndLadders([][]int{
                {-1,-1,19,10,-1},
                {2,-1,-1,6,-1},
                {-1,17,-1,19,-1},
                {25,-1,20,-1,-1},
                {-1,-1,-1,-1,15}})) //2
}

func snakesAndLadders(board [][]int) int {
        /*
        As we are looking for a shortest path, a breadth-first search is ideal.
        BFS 解决最短路径问题...
        BFS 和DFS 不同 BFS 是正向考虑问题的.dist记录到开始点的最近路径距离
        */
        // bfs
        n := len(board)
        queue := make([]int,0)
        // 一方面记录结果,另一方面用于访问元素去重复.
        // key是顶点,value表示该顶点距离起点的最小距离. bfs,从前往后保证最短路径. dist用来记录结果.
        dist := make(map[int]int)
        // 初始化,从1出发
        dist[1]=0
        queue = append(queue,1)
        for len(queue) > 0 {
                s := queue[0]  // 当前访问位置s
                queue = queue[1:]
                if s >= n*n {
                        return dist[s]
                }
                for s2:=s+1;s2<=s+6&&s2<=n*n;s2++ {
                        //下一个可能位置s2.
                        // 两种可能,如果是-1,则本身位置是下一个
                        // 否则就是board中的值,涉及到坐标变化.trans
                        i, j := trans(s2, n)
                        ns := board[i][j]
                        if board[i][j] == -1 {
                                ns = s2
                        }
                        // 如果目的节点ns没访问过,则他肯定最小是dist[s]+1,并且放到队列做bfs
                        if _,ok := dist[ns];!ok {
                                dist[ns]=dist[s]+1
                                queue = append(queue,ns)
                        }
                }
        }
        return -1
}
func trans(start int,n int) (int,int) {
        i := n-(start-1)/n-1
        j := (start-1)%n
        if (n-i) % 2 == 0 {
                j = n-1-(start-1)%n
        }
        return i,j
}
//////////////////////////////////////////////////
func snakesAndLadders1(board [][]int) int {
        cache := make([]int,(len(board))*(len(board))+1)
        min := solve(1,board,len(board),&cache)
        if min < 0 {
                return -1
        }
        return min
}
func solve(start int, board [][]int, n int,cache *[]int) int {
        min := math.MaxInt64
        for index:=start+1;index<=start+6;index++ {
                if index >= n*n {
                        (*cache)[start]=1
                        return 1
                }
                i,j := trans(index,n)
                nstart := index
                if board[i][j] != -1 {
                        nstart=board[i][j]
                }
                if nstart >= n*n {
                        (*cache)[start]=1
                        return 1
                }
                if (*cache)[nstart] == 0 {
                        (*cache)[nstart]=-1
                        r := solve(nstart, board, n, cache)
                        if r > 0 && min > r {
                                min = r
                        }
                } else {
                        r := (*cache)[nstart]
                        if r>0 && min > r {min=r}

                }
        }
        if min == math.MaxInt64 {
                (*cache)[start]=-1
                return -1
        }
        (*cache)[start]=min+1
        return min+1
}

