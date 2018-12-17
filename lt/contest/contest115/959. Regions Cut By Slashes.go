package main

import "fmt"

/*
In a N x N grid composed of 1 x 1 squares, each 1 x 1 square consists of a /, \, or blank space.  These characters divide the square into contiguous regions.

(Note that backslash characters are escaped, so a \ is represented as "\\".)

Return the number of regions.



Example 1:

Input:
[
  " /",
  "/ "
]
Output: 2
Explanation: The 2x2 grid is as follows:

Example 2:

Input:
[
  " /",
  "  "
]
Output: 1
Explanation: The 2x2 grid is as follows:

Example 3:

Input:
[
  "\\/",
  "/\\"
]
Output: 4
Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
The 2x2 grid is as follows:

Example 4:

Input:
[
  "/\\",
  "\\/"
]
Output: 5
Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
The 2x2 grid is as follows:

Example 5:

Input:
[
  "//",
  "/ "
]
Output: 3
Explanation: The 2x2 grid is as follows:



Note:

1 <= grid.length == grid[0].length <= 30
grid[i][j] is either '/', '\', or ' '.

*/

func main() {
    fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))//5
    fmt.Println(regionsBySlashes([]string{  " /", "  "}))//1
    fmt.Println(regionsBySlashes([]string{  "\\/", "/\\"}))//4
    fmt.Println(regionsBySlashes([]string{  "//", "/ "}))//3
    fmt.Println(regionsBySlashes([]string{" /", "/ "}))//2

}
func regionsBySlashes(grid []string) int {
    m,n := len(grid)*3,len(grid[0])*3
    visit := make([][]bool,0)
    for i:=0;i<m;i++ {
        visit = append(visit,make([]bool,n))
    }
    ngrid := make([][]byte,0)
    for i:=0;i<m;i++ {
        ngrid = append(ngrid,make([]byte,n))
    }
    /*
    放大两倍是不行的， 处理不了
    [
    "//",
    "/ "]
    这种case，所以，放大到三倍
    */
    for i:=0;i<m/3;i++ {
        for j:=0; j<n/3; j++ {
            if grid[i][j]=='\\' {
                ngrid[i*3][j*3],ngrid[i*3][j*3+1],ngrid[i*3][j*3+2]='1','0','0'
                ngrid[i*3+1][j*3],ngrid[i*3+1][j*3+1],ngrid[i*3+1][j*3+2]='0','1','0'
                ngrid[i*3+2][j*3],ngrid[i*3+2][j*3+1],ngrid[i*3+2][j*3+2]='0','0','1'
            } else if grid[i][j]=='/' {
                ngrid[i*3][j*3],ngrid[i*3][j*3+1],ngrid[i*3][j*3+2]='0','0','1'
                ngrid[i*3+1][j*3],ngrid[i*3+1][j*3+1],ngrid[i*3+1][j*3+2]='0','1','0'
                ngrid[i*3+2][j*3],ngrid[i*3+2][j*3+1],ngrid[i*3+2][j*3+2]='1','0','0'
            } else {
                ngrid[i*3][j*3+0],ngrid[i*3][j*3+1],ngrid[i*3][j*3+2]='0','0','0'
                ngrid[i*3+1][j*3+0],ngrid[i*3+1][j*3+1],ngrid[i*3+1][j*3+2]='0','0','0'
                ngrid[i*3+2][j*3+0],ngrid[i*3+2][j*3+1],ngrid[i*3+2][j*3+2]='0','0','0'
            }
        }
    }
    //for i:=0;i<m;i++ {
        //fmt.Println(string(ngrid[i]))
        //    fmt.Println(visit[i])
    //}
    ans := 0
    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            if ngrid[i][j]=='0' && visit[i][j]==false{
                dfs(i,j,m,n,ngrid,&visit)
                ans += 1
            }
        }
    }
    return ans
}
func dfs(i,j,m,n int,grid [][]byte,visit *[][]bool) {
    (*visit)[i][j]=true
    dir := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    for _,d := range dir {
        u,v := i+d[0],j+d[1]
        if u<0||u>=m||v<0||v>=n||grid[u][v]!='0'||(*visit)[u][v] {continue}
        dfs(u,v,m,n,grid,visit)
    }

}