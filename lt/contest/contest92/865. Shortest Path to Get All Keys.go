package main

import "fmt"
/*
We are given a 2-dimensional grid. "." is an empty cell, "#" is a wall, "@" is the starting point, ("a", "b", ...) are keys, and ("A", "B", ...) are locks.

We start at the starting point, and one move consists of walking one space in one of the 4 cardinal directions.  We cannot walk outside the grid, or walk into a wall.  If we walk over a key, we pick it up.  We can't walk over a lock unless we have the corresponding key.

For some 1 <= K <= 6, there is exactly one lowercase and one uppercase letter of the first K letters of the alphabet in the grid.  This means that there is exactly one key for each lock, and one lock for each key; and also that the letters used to represent the keys and locks were chosen in the same order as the alphabet.

Return the lowest number of moves to acquire all keys.  If it's impossible, return -1.



Example 1:

Input: ["@.a.#","###.#","b.A.B"]
Output: 8
Example 2:

Input: ["@..aA","..B#.","....b"]
Output: 6


Note:

1 <= grid.length <= 30
1 <= grid[0].length <= 30
grid[i][j] contains only '.', '#', '@', 'a'-'f' and 'A'-'F'
grid[i] is a string of length K
The number of keys is in [1, 6].  Each key has a different letter and opens exactly one lock.
*/

func shortestPathAllKeys(grid []string) int {
        if len(grid)<=0 {return -1}
        g := make([][]byte,0)
        y := make([][]bool,0)
        v := make(map[byte]bool)
        si,sj := 0,0
        m,n := len(grid),len(grid[0])
        num := 0
        for i:=0;i<len(grid);i++ {
                g = append(g,[]byte(grid[i]))
                y = append(y,make([]bool,len(grid[i])))
                for j:=0;j<len(grid[i]);j++ {
                        if grid[i][j] >= 'a' &&grid[i][j] <='z' {
                                num += 1
                        }
                        if grid[i][j]=='@' {
                                si,sj = i,j
                        }
                }
        }
        dep := m*n
        dfs(g,y,si,sj,m,n,v,0,&dep,num*2)
        if dep == m*n {return -1}
        return dep
}
var direction [][]int = [][]int{
        []int{0,1},
        []int{0,-1},
        []int{1,0},
        []int{-1,0},
}
func dfs(grid [][]byte,visit [][]bool,i,j int,m,n int, az map[byte]bool,dep int,step *int,lens int)  {
        fmt.Println(i,j,dep,az,m,n,lens,visit)
        if ((grid[i][j] >= 'a' &&grid[i][j] <='z')||(grid[i][j] >= 'A' &&grid[i][j] <='Z')) && len(az) == lens-1 {
                fmt.Println("XXX",i,j,az,dep,lens)
                if *step > dep {
                        *step=dep
                }
                return
        }
        if grid[i][j] >= 'a' &&grid[i][j] <='z' {
                az[grid[i][j]]=true
        }
        if grid[i][j] >= 'A' &&grid[i][j] <='Z' {
                az[grid[i][j]]=true
        }
        for _,d := range direction {
                ni,nj := i+d[0],j+d[1]
                if ni < 0 || nj < 0 || ni >=m || nj >=n || grid[i][j] == '#' || grid[ni][nj]=='@' || visit[ni][nj]==true{
                        continue
                }
                if grid[ni][nj] >= 'A' &&grid[ni][nj] <='Z' {
                        if  _,ok := az[grid[ni][nj]-'A'+'a'];ok {
                                if az[grid[ni][nj]-'A'+'a']==false{
                                        continue
                                }
                        }
                }
                fmt.Println(ni,nj)
                if grid[ni][nj] >= 'a' &&grid[ni][nj] <='z' {
                        if _,ok := az[grid[ni][nj]-'a'+'A'];ok{
                                if az[grid[ni][nj]-'a'+'A']==false {
                                        continue
                                }
                        }
                }
                visit[ni][nj]=true
                dfs(grid,visit,ni,nj,m,n,az,dep+1,step,lens)
                visit[ni][nj]=false
        }
}
func main() {
        //fmt.Println(shortestPathAllKeys([]string{
        //        "@.a.#","###.#","b.A.B",
        //}))
        fmt.Println(shortestPathAllKeys([]string{
//                "@..aA","..B#.","....b",
                "@...a",".###A","b.BCc",
        }))
}
