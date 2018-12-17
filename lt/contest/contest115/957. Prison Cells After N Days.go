package main

import "fmt"

/*
There are 8 prison cells in a row, and each cell is either occupied or vacant.

Each day, whether the cell is occupied or vacant changes according to the following rules:

If a cell has two adjacent neighbors that are both occupied or both vacant,
then the cell becomes occupied.
Otherwise, it becomes vacant.

(Note that because the prison is a row, the first and the last cells
in the row can't have two adjacent neighbors.)

We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)



Example 1:

Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation:
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

Example 2:

Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]


Note:

cells.length == 8
cells[i] is in {0, 1}
1 <= N <= 10^9
*/

func main() {
    fmt.Println(prisonAfterNDays([]int{0,1,0,1,1,0,0,1},32))
    fmt.Println(prisonAfterNDays([]int{1,0,0,1,0,0,1,0},1000000000))

}
/*
8个int，最多32个,2^8=32
变化过程是不可逆的，所以可能出现多个dayi情况都会到day i+1
所有答案分两个过程
1.找到循环 取模
2.把剩下的day变化simulate完
*/
func prisonAfterNDays(cells []int, N int) []int {
    dict := make(map[[8]int]int)
    key := [8]int{}
    for i:=0;i<8;i++ {key[i]=cells[i]}
    idx := 1
    for N > 0 {
        if _,ok := dict[key];ok {
            N = N%(idx-dict[key])
            break
        } else {
            dict[key]=idx
            key = trans(key)
            idx += 1
            N-=1
        }
    }
    // 取模之后的day
    for N > 0 {
        key = trans(key)
        N -= 1
    }
    ans := make([]int,8)
    for i:=0;i<8;i++ {ans[i]=key[i]}
    return ans
}
func trans(a [8]int) [8]int {
    b := [8]int{}
    for i:=1;i<7;i++ {
        if a[i-1]==0&&a[i+1]==0 || a[i-1]==1&&a[i+1]==1 {
            b[i]=1
        }
    }
    return b
}