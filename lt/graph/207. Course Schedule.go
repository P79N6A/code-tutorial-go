package main

import "fmt"

/*
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

Example 1:

Input: 2, [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: 2, [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0, and to take course 0 you should
             also have finished course 1. So it is impossible.
Note:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
*/
func main() {
    fmt.Println(canFinish(2,[][]int{{1,0}}))
    fmt.Println(canFinish(2,[][]int{{1,0},{0,1}}))
}
// build graph and check circle.
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int,numCourses)
    for _,p := range prerequisites {
        if len(graph[p[0]]) <= 0 {
            graph[p[0]] = []int{p[1]}
        } else {
            graph[p[0]] = append(graph[p[0]],p[1])
        }
    }
    visit := make([]bool,numCourses)
    for i:=0;i<numCourses;i++ {
        if visit[i]==true {
            continue
        }
        if dfs(i,graph,&visit) {
            return false
        }
    }
    return true
}

func dfs(v int, graph [][]int,visit *[]bool) bool {
    (*visit)[v]=true
    for _,u := range graph[v] {
        if (*visit)[u] == true {return true}
        if dfs(u,graph,visit) {
            return true
        }
    }
    (*visit)[v]=false
    return false
}
