package main

import "fmt"

/*
给出无向图G，给出m中颜色，问让图G中任意相邻的节点颜色都不同

问题1： 用m中颜色是否能做到上述要求？
问题2： 如果做到，有多少种染色方法
问题3： 如果没有给出颜色有多少种，问最少的能够满足染色的颜色m的个数？

思路：
图如果点少变多就用邻接矩阵，否则就邻接表
使用回溯的方法。染色从1到m，判断相邻点是否可以染指定颜色，不行就换，可以就继续。
成功完成一次涂色就记录一次结果。


示例graph
                +---+
             XXX| 1 |XXXXX
           XXX  +-+-+     XXXX
         XX       |          XX
      XXXX        |            X
     XX           |            XX
 +----+        +--+--+          XX
 | 2  +--------+  3  |           X
 +-+--X        +--+--+           X
   |   XXX        |              X
   |     XXX      |             XX
   |        XXXX  |             X
+--+--+        X--+--+       XXX
|  5  +--------+  4  |XXXXXXXX
+-----+        +-----+

参考 ： https://www.geeksforgeeks.org/backttracking-set-5-m-coloring-problem/

*/

var graph = [][]int{
    {0, 0, 0, 0, 0, 0},
    {0, 0, 1, 1, 1, 0},
    {0, 1, 0, 1, 1, 1},
    {0, 1, 1, 0, 1, 0},
    {0, 1, 1, 1, 0, 1},
    {0, 0, 1, 0, 1, 0},
}

func main() {
    counter := 0
    m := 5
    color := make([]int, m+1)

    /*
    如果节点染色1，2不一样，则不固定节点1的染色。
    colorDFS(graph,m,1,4, &color, &counter)
    */
    //如果节点染成什么颜色不关心，只关注分多少组，则固定第一个节点颜色即可。
    color[1]=1
    colorDFS(graph, m, 2,4, &color, &counter)
    fmt.Println(counter)
    // 判断是否可以完成染色
    color = make([]int, m+1)
    fmt.Println(canColor(graph,m,1,3,&color))
    fmt.Println(canColor1(graph,m,1,3,color))
    color = make([]int, m+1)
    fmt.Println(canColor(graph,m,1,4,&color))
    fmt.Println(canColor1(graph,m,1,4,color))

    fmt.Println(minColor(graph,m))
}

func  minColor(g [][]int, nodeNum int) int {
    /*
    计算最少的染色个数，NP问题
    贪心的算法,并非最优解，但可以计算上界upperbound
    Basic Greedy Coloring Algorithm:

    1. Color first vertex with first color.
    2. Do following for remaining V-1 vertices.
    ….. a) Consider the currently picked vertex and color it with the
    lowest numbered color that has not been used on any previously
    colored vertices adjacent to it. If all previously used colors
    appear on vertices adjacent to v, assign a new color to it.
    */
    colorSeq := make([]int,nodeNum+1)
    color := 1
    for i:=1;i<=nodeNum;i++ {
        for j:=1;j<=color;j++ {
            colorSeq[i]=j
            if isValid(g,nodeNum,i,colorSeq) {
                break
            }
            if j == color {
                // all color can not valid.
                // increase color number
                color += 1
                colorSeq[i] = color
            }
        }
    }
    return color
}

func canColor1(g [][]int,nodeNum,id,colorNum int,color []int) bool {

    for i:=1;i<=nodeNum;i++ {
        for j:=1;j<=colorNum;j++ {
            color[i]=j
            if isValid(g,nodeNum,i,color) {
                break
            }
            if j==colorNum {
                //所有的染色尝试都不行了，则返回false
                return false
            }
        }
    }
    return true
}
func canColor(g [][]int,nodeNum,id,colorNum int,color *[]int) bool {
    if id > nodeNum {
        return true
    }
    // 只判断是否能完成染色，不关心多少种方案
    for i:=1;i<=colorNum;i++ {
        (*color)[id] = i
        if isValid(g, nodeNum, id, *color) {
            if canColor(g, nodeNum, id+1, colorNum, color) {
                return true
            }
        }
    }
    return false
}

func colorDFS(g [][]int, nodeNum, id, colorNum int, color *[]int, count *int) {
    // 功能：在graph上，基于给定的颜色个数colorNum，问有多少种染色方案count,如果count为0，则无法完成染色
    // 输入参数  graph【邻接矩阵】, 节点个数， 当前处理节点， 颜色个数， 染色方案， 染色方案数
    if id > nodeNum { // 完成了一次染色
        *count += 1
        return
    }
    // 回溯过程，尝试各个染色方案，
    for i := 1; i <= colorNum; i++ {
        (*color)[id] = i
        if isValid(g, nodeNum, id, *color) {
            colorDFS(g, nodeNum, id+1, colorNum, color, count)
        }
        (*color)[id] = 0
    }
}
func isValid(g [][]int, nodeNum, id int, color []int) bool {
    // 看如果id染了指定的颜色color[id],那么是否会产生冲突
    for i := 1; i <= nodeNum; i++ {
        if g[id][i] == 1 { // have edge
            if color[id] == color[i] { //相邻的有相同颜色了
                return false
            }
        }
    }
    return true
}
