package main

import "fmt"

/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.
*/
func isValidSudoku(board [][]byte) bool {
    /*
    使用一个hash即可
    line-0-9 key 表示，第0行，出现了9
    row-1-8 key 表示  第一列 出现了8
    sub-0-1-7 key表示，第0，1个小方块里面出现了7
    */
    valid := make(map[string]bool)
    // key :line-1,row-1,s1-1-1
    for i:=0;i<len(board);i++ {
        for j:=0;j<len(board[0]);j++ {
            if board[i][j]=='.'{continue}
            line := fmt.Sprintf("line-%d-%v",i,board[i][j])
            if valid[line]==true{return false}
            valid[line]=true
            row := fmt.Sprintf("row-%d-%v",j,board[i][j])
            if valid[row]==true{return false}
            valid[row]=true
            sub := fmt.Sprintf("sub-%d-%d-%v",i/3,j/3,board[i][j])
            if valid[sub]==true{return false}
            valid[sub]=true
        }
    }
    return true
}
func main() {
    
}
