package main

import "fmt"

/*
A Tic-Tac-Toe board is given as a string array board. Return True if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.

The board is a 3 x 3 array, and consists of characters " ", "X", and "O".  The " " character represents an empty square.

Here are the rules of Tic-Tac-Toe:

Players take turns placing characters into empty squares (" ").
The first player always places "X" characters, while the second player always places "O" characters.
"X" and "O" characters are always placed into empty squares, never filled ones.
The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Example 1:
Input: board = ["O  ", "   ", "   "]
Output: false
Explanation: The first player always plays "X".

Example 2:
Input: board = ["XOX", " X ", "   "]
Output: false
Explanation: Players take turns making moves.

Example 3:
Input: board = ["XXX", "   ", "OOO"]
Output: false

Example 4:
Input: board = ["XOX", "O O", "XOX"]
Output: true
Note:

board is a length-3 array of strings, where each string board[i] has length 3.
Each board[i][j] is a character in the set {" ", "X", "O"}.

*/

func main() {
    fmt.Println(validTicTacToe([]string{"XOX"," X ","   "}))
}
func validTicTacToe(board []string) bool {
    /*
    题目问验证是不是一个有效的游戏结果。
    什么是有效的？(X先下)
    1.没有结束的中间结果，X个数等于O或者等于O+1
    2.当X已经成了三联了，则X=O+1
    3.当O已经成了三联了，则X=O
    */
    X,O := 0,0
    for i:=0;i<len(board);i++ {
        for j:=0;j<len(board);j++ {
            if board[i][j]=='X' {X+=1}
            if board[i][j]=='O' {O+=1}
        }
    }
    if finish(board,'X') {
        return X == O+1
    }
    if finish(board,'O') {
        return X==O
    }
    return X==O+1 || X==O
}
func finish(board []string, t byte) bool {
    for i:=0;i<len(board);i++ {
        f := true
        for j:=0;j<len(board[i]);j++ {
            if board[i][j]!=t {
                f = false
                break
            }
        }
        if f {return true}
    }
    for i:=0;i<len(board);i++ {
        f := true
        for j:=0;j<len(board[i]);j++ {
            if board[j][i]!=t {
                f = false
                break
            }
        }
        if f {return true}
    }
    if board[0][0] == t && board[1][1] == t && board[2][2] == t {return true}
    if board[0][2] == t && board[1][1] == t && board[2][0] == t {return true}
    return false
}