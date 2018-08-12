package main
/*
Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return true if and only if it is possible to split everyone into two groups in this way.



Example 1:

Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
Example 2:

Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Example 3:

Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false


Note:

1 <= N <= 2000
0 <= dislikes.length <= 10000
1 <= dislikes[i][j] <= N
dislikes[i][0] < dislikes[i][1]
There does not exist i != j for which dislikes[i] == dislikes[j].

*/
func main() {
}

func possibleBipartition(N int, dislikes [][]int) bool {
        A := make(map[int]bool)
        B := make(map[int]bool)
        return solve(A,B,dislikes,0)
}

func solve(A,B map[int]bool,dis [][]int, idx int) bool {
        if idx >= len(dis) {
                return true
        }
        okA0:=A[dis[idx][0]]
        okB1:=B[dis[idx][1]]

        okA1:=A[dis[idx][1]]
        okB0:=B[dis[idx][0]]

        if okA0 {

        } else {

        }

        if !okA0 && !okB1 {
                A[dis[idx][0]]=true
                B[dis[idx][1]]=true
                if !solve(A,B,dis,idx+1) {
                        return false
                }
                A[dis[idx][0]]=false
                B[dis[idx][1]]=false
        }
        if !okA1 && !okB0 {
                A[dis[idx][1]]=true
                B[dis[idx][0]]=true
                if !solve(A,B,dis,idx+1) {
                        return false
                }
                A[dis[idx][1]]=false
                B[dis[idx][0]]=false
        }
        return false
}