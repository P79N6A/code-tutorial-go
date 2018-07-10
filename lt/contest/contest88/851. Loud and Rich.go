package main

import (
        "fmt"
)

/*
851. Loud and Rich
User Accepted: 138
User Tried: 195
Total Accepted: 138
Total Submissions: 260
Difficulty: Medium
In a group of N people (labelled 0, 1, 2, ..., N-1), each person has different amounts of money,
 and different levels of quietness.

For convenience, we'll call the person with label x, simply "person x".

We'll say that richer[i] = [x, y] if person x definitely has more money than person y.
Note that richer may only be a subset of valid observations.

Also, we'll say quiet[x] = q if person x has quietness q.

Now, return answer, where answer[x] = y if y is the least quiet person (that is,
the person y with the smallest value of quiet[y]),
among all people who definitely have equal to or more money than person x.



Example 1:

Input: richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
Output: [5,5,2,5,4,5,6,7]
Explanation:
answer[0] = 5.
Person 5 has more money than 3, which has more money than 1, which has more money than 0.
The only person who is quieter (has lower quiet[x]) is person 7, but
it isn't clear if they have more money than person 0.

answer[7] = 7.
There isn't anyone who definitely has more money than person 7, so the person who definitely has
equal to or more money than person 7 is just person 7.

The other answers can be filled out with similar reasoning.
Note:

1 <= quiet.length = N <= 500
0 <= quiet[i] < N, all quiet[i] are different.
0 <= richer.length <= N * (N-1) / 2
0 <= richer[i][j] < N
richer[i][0] != richer[i][1]
richer[i]'s are all different.
The observations in richer are all logically consistent.

*/
func loudAndRich(richer [][]int, quiet []int) []int {
        more := make([][]int,len(quiet))
        for _,r := range richer {
                if len(more[r[1]]) <= 0 {
                        more[r[1]] = make([]int,0)
                }
                more[r[1]] = append(more[r[1]],r[0])
        }
        res := make([]int,len(quiet))
        for i:=0;i<len(quiet);i++ {
                dfs2(more,quiet,i,&res)
                //res = append(res,r)
        }
        return res
}
func dfs2(more [][]int,quiet []int,i int, ret *[]int) int {
        if (*ret)[i] > 0 {
                return (*ret)[i]
        }
        if len(more[i]) <= 0 {
                (*ret)[i] = i
                return i
        }
        min := quiet[i]
        min_i := i
        for _,v := range more[i] {
                vi := dfs2(more,quiet,v,ret)
                if min > quiet[vi] {
                        min = quiet[vi]
                        min_i = vi
                }
        }
        (*ret)[i] = min_i
        return min_i
}
func dfs(more [][]int,quiet []int,i int) int {
        if len(more[i]) <= 0 {
                return i
        }
        min := quiet[i]
        min_i := i
        for _,v := range more[i] {
                vi :=  dfs(more,quiet,v)
                if min > quiet[vi] {
                        min = quiet[vi]
                        min_i=vi
                }
        }
        return min_i
}
func main() {
        //*
        fmt.Println(loudAndRich([][]int{
                []int{3,7},
                []int{4,3},
                []int{1,0},
                []int{2,1},
                []int{3,1},
                []int{5,3},
                []int{6,3},
        },[]int{3,2,5,4,6,1,7,0}))
        //*/
        fmt.Println(loudAndRich([][]int{
                []int{0,1},
                []int{1,2},
        },[]int{0,1,2}))
}
