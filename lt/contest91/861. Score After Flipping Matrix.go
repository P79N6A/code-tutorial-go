package main
/*
We have a two dimensional matrix A where each value is 0 or 1.

A move consists of choosing any row or column, and toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s.

After making any number of moves, every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.

Return the highest possible score.



Example 1:

Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
Output: 39
Explanation:
Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39


最终结果是加和,则第一位是1才最重要,保证是1了之后,看是否还能比之前的大,就flip column
*/

import "fmt"

func matrixScore(A [][]int) int {
        /*
        greedy, 第一列都变成1.【通过行变换】
        其他列，看0多还是1多，如果0多，转换一下。
        */
        m := len(A)
        if m <= 0 {
                return 0
        }
        n := len(A[0])
        for i := 0; i < m; i++ {
                if A[i][0] == 0 {
                        A = flipRow(A, m, n, i)
                }
        }
        for j := 1; j < n; j++ {
                z, o := columnZeroOne(A, m, n, j)
                if z > o {
                        A = flipColumn(A, m, n, j)
                }
        }
        ret := 0
        for i := 0; i < m; i++ {
                ret += calRow(A, m, n, i)
        }
        return ret
}

func columnZeroOne(A [][]int, m, n int, column int) (int, int) {
        zero := 0
        for i := 0; i < m; i++ {
                if A[i][column] == 0 {
                        zero += 1

                }
        }
        return zero, m - zero
}
func calRow(A [][]int, m, n int, row int) int {
        ret := 0
        for i := n - 1; i >= 0; i-- {
                if A[row][i] == 1 {
                        ret += 1
                }
                ret <<= 1
        }
        return ret
}
func flipRow(A [][]int, m, n int, row int) [][]int {
        for i := 0; i < n; i++ {
                if A[row][i] == 0 {
                        A[row][i] = 1
                } else {
                        A[row][i] = 0
                }
        }
        return A
}
func flipColumn(A [][]int, m, n int, column int) [][]int {
        for i := 0; i < m; i++ {
                if A[i][column] == 0 {
                        A[i][column] = 1
                } else {
                        A[i][column] = 0
                }
        }
        return A
}

func main() {
        fmt.Println(matrixScore([][]int{
                []int{0,0,1,1},
                []int{1,0,1,0},
                []int{1,1,0,0},
        }))
}
