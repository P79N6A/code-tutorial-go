package main

import (
        "fmt"
    "sort"
)
/*
A string S of lowercase letters is given.  Then, we may make any number of moves.

In each move, we choose one of the first K letters (starting from the left), remove it, and place it at the end of the string.

Return the lexicographically smallest string we could have after any number of moves.



Example 1:

Input: S = "cba", K = 1
Output: "acb"
Explanation:
In the first move, we move the 1st character ("c") to the end, obtaining the string "bac".
In the second move, we move the 1st character ("b") to the end, obtaining the final result "acb".
Example 2:

Input: S = "baaca", K = 3
Output: "aaabc"
Explanation:
In the first move, we move the 1st character ("b") to the end, obtaining the string "aacab".
In the second move, we move the 3rd character ("c") to the end, obtaining the final result "aaabc".


Note:

1 <= K <= S.length <= 1000
S consists of lowercase letters only.
*/
func main() {
        fmt.Println(orderlyQueue("bxxweesaca",1))
}
/*
问题：给了一个字符串，和一个长度K。规定一个操作可以将前K个字符的某一个，转移到最后。这个操作可以做任意次。 问最终转换出来的最小的字符串是什么？
给了两个例子，K=1,  “cba”=>”acb“   K=2, “baaca”=>”aaabc”
思路：找了个case推演了一下，想找规律，前K个肯定可以是最小的字符，但后边的如何？转了几次，如果每次讲不在最小的K个中最小的转到后边去，则后边可以是最小的，但结果可能不对。 假设了一下是不是排个序就是最终结果了，因为排序后的结果理论最小。但当K=1的时候又不是。
换个思路这个序列 A1A2….Ai-1AiAi+1….An  假如K>1的，则通过转换，可以得到这样的结果
A1A2….AiAi+1….An=>AiAi+1….AnA1A2….Ai-1  将A1-Ai-1挪到最后边去按顺序。
因为K>1，则吧Ai+1先挪走 =>AiAi+2….AnA1A2….Ai-1Ai+1
再把Ai挪走 =>Ai+2….AnA1A2….Ai-1Ai+1Ai
再把Ai+1后边的挪走 =>A1A2….Ai-1Ai+1AiAi+2….An
会发现，这个序列跟第一个序列的区别就是 Ai Ai+1做了一次swap，这跟冒泡排序是一样的，并且可以推广到一般情况，任意两个相邻的都可以交换。那么只要K>1则最终最小的就是排序结果【冒泡排序】
K=1的时候呢？ 这个是没法做swap的，因为规定了往后移动的只能是第一个。

*/
func orderlyQueue(S string, K int) string {
    if K > 1 {
        bs := []byte(S)
        sort.Slice(bs, func(i, j int) bool {
            return bs[i] < bs[j]
        })
        return string(bs)
    }
    min := S
    S = S+S // 找到所有的变化后子串，循环子串，得到最小
    for i:=0;i<len(min);i++ {
        if S[i:i+len(min)] < min {
            min = S[i:i+len(min)]
        }
    }
    return min
}
