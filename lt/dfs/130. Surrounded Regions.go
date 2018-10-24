package main

import "fmt"

/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/


func main() {
    b := make([][]byte,0)
    b = append(b,[]byte("XXXX"))
    b = append(b,[]byte("XOOX"))
    b = append(b,[]byte("XXOX"))
    b = append(b,[]byte("XOXX"))
    solve(b)
    for _,bb := range b {
        fmt.Println(string(bb))
    }

}
/*
问题：将X包裹的部分都设置成X。如何识别哪个区域被X包裹了？
里面不行，就是边缘处理，将X的非包裹区域设置成A,剩下的O就是被包裹的了，再把A设置回来。
*/
func solve(board [][]byte)  {
    if len(board)<=0{return}
    m,n := len(board),len(board[0])
    for i:=0;i<m;i++ {dfs(i,0,&board,m,n)}//左边缘
    for i:=0;i<m;i++ {dfs(i,n-1,&board,m,n)}//右边缘
    for i:=0;i<n;i++ {dfs(0,i,&board,m,n)}//上边缘
    for i:=0;i<n;i++ {dfs(m-1,i,&board,m,n)}//下边缘
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            if board[i][j]=='O' {board[i][j]='X'}
            if board[i][j]=='A' {board[i][j]='O'}
        }
    }
}
var direction = [][]int{{0,1},{1,0},{0,-1},{-1,0}}
func dfs(ui,uj int,b *[][]byte,m,n int) { //遍历，X是边界，将O设置成A;
    if (*b)[ui][uj]!='O' {return}
    (*b)[ui][uj]='A'
    for _,d := range direction {
        i,j := ui+d[0],uj+d[1]
        if i<0||i>=m ||j<0 || j>=n||(*b)[i][j]!='O' {
            continue
        }
        dfs(i,j,b,m,n)
    }
}