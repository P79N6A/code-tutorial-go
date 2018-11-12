package main

import "fmt"
/*
Given a string S, count the number of distinct, non-empty subsequences of S .

Since the result may be large, return the answer modulo 10^9 + 7.



Example 1:

Input: "abc"
Output: 7
Explanation: The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".
Example 2:

Input: "aba"
Output: 6
Explanation: The 6 distinct subsequences are "a", "b", "ab", "ba", "aa" and "aba".
Example 3:

Input: "aaa"
Output: 3
Explanation: The 3 distinct subsequences are "a", "aa" and "aaa".




Note:

S contains only lowercase letters.
1 <= S.length <= 2000

*/

func main() {
    fmt.Println(distinctSubseqII("abc"))
    fmt.Println(distinctSubseqII("aba"))
    fmt.Println(distinctSubseqII("aaa"))
}
/*
如何思考?问题的关键在子集的处理上是如何有叠加效果的?如何继承之前的结果.
子集是有顺序相关性的.记录下以任意字符结尾的子集个数.每次计算更新结果
*/
func distinctSubseqII(S string) int {
    mod := int(1e9)+7
    endwith := make([]int,26)//
    for _,s := range S {
        /*
        endwith 记录以某个字符结尾的所有非空子集个数
        新来一个字符,可以追加到以任意字符结尾的非空子集上.那就是把之前的结果加和了赋给自己.
        */
        endwith[s-'a']=sum(endwith)%mod+1
    }
    return sum(endwith)%mod
}
func sum(rec []int) int {
    ans := 0
    for _,r := range rec {
        ans += r
    }
    return ans
}