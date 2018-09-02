package main

import (
        "strings"
        "sort"
        "fmt"
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
        fmt.Println(orderlyQueue("bxxweesaca",3))
}


type SortB []byte

func (s SortB) Len() int {return len(s)}

func (s SortB) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
}

func (s SortB) Less(i, j int) bool {
        return (s[i]) < (s[j])
}

func orderlyQueue(S string, K int) string {
        bs := []byte(S)
        sort.Sort(SortB(bs))
        finish := string(bs)[:K]
        fset := make(map[byte]bool)
        for i:=0;i<len(finish);i++ {
                fset[finish[i]]=true
        }
        return solve(finish,fset,K,S)
}
func solve(finish string,fset map[byte]bool,K int, S string) string {
        //fmt.Println("XXX",S)
        if strings.HasPrefix(S,finish) {
                return S
        }
        i := 0
        for ;i<K;i++{
                if finish[i] != S[i] {
                        break
                }
        }
        mi := 0
        var min byte = 'z'+1
        for ;i<K;i++ {
                if fset[S[i]] {
                        continue
                }
                if S[i] < min {
                        min = S[i]
                        mi = i
                }
        }
        nS := S[:mi] + S[(mi+1):] + string(S[mi])
        return solve(finish,fset,K,nS)
}