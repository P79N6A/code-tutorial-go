package main

import (
    "sort"
    "fmt"
)

/*

You have a number of envelopes with widths and heights given as a pair of integers (w, h). One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

What is the maximum number of envelopes can you Russian doll? (put one inside other)

Example:
Given envelopes = [[5,4],[6,4],[6,7],[2,3]], the maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
*/
type sortenv [][]int
func (s sortenv)Len() int {
    return len(s)
}
func (s sortenv)Less(i,j int) bool {
    if s[i][0]==s[j][0] {
        return s[i][1]<s[j][1]
    }
    return s[i][0]<s[j][0]
}
func (s sortenv)Swap(i,j int) {
    s[i],s[j]=s[j],s[i]
}

func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes) <= 0 {return 0}
    dp := make([]int,len(envelopes))
    for i:=0;i<len(dp);i++ {dp[i]=1}
    sort.Sort(sortenv(envelopes))
    max := 1
    for i:=1;i<len(envelopes);i++ {
        for k:=0;k<i;k++ {
            if envelopes[i][0]>envelopes[k][0] && envelopes[i][1]>envelopes[k][1] {
                if dp[i] < dp[k]+1 {
                    dp[i]=dp[k]+1
                }
            }
        }
        if max < dp[i] {
            max = dp[i]
        }
    }
    return max
}
func main() {
    fmt.Println(maxEnvelopes([][]int{
        []int{1,1},
    }))
    fmt.Println(maxEnvelopes([][]int{
        //[[10,8],[1,12],[6,15],[2,18]]
        []int{10,8},
        []int{1,12},
        []int{6,15},
        []int{2,18},
    }))
    
}
