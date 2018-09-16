package main

import "fmt"

/*
We are given S, a length n string of characters from the set {'D', 'I'}. (These letters stand for "decreasing" and "increasing".)

A valid permutation is a permutation P[0], P[1], ..., P[n] of integers {0, 1, ..., n}, such that for all i:

If S[i] == 'D', then P[i] > P[i+1], and;
If S[i] == 'I', then P[i] < P[i+1].
How many valid permutations are there?  Since the answer may be large, return your answer modulo 10^9 + 7.



Example 1:

Input: "DID"
Output: 5
Explanation:
The 5 valid permutations of (0, 1, 2, 3) are:
(1, 0, 3, 2)
(2, 0, 3, 1)
(2, 1, 3, 0)
(3, 0, 2, 1)
(3, 1, 2, 0)


Note:

1 <= S.length <= 200
S consists only of characters from the set {'D', 'I'}.
*/

func main() {
    fmt.Println(numPermsDISequence("DID"))
    fmt.Println(numPermsDISequence("IDDDIIDIIIIIIIIDIDID")) // 853197538
    fmt.Println(numPermsDISequence("IIDIIDDIDDDDIIDDIDIDIDDDDIDDDIIIIDDIDDDDIDIIDDIDID")) // 997381513
    //642107237
    fmt.Println(numPermsDISequence("IIDDIDDIIDDIDIDDIDDDDIIIDIDIDDDDDIIDIDDIDIIDIDDIIIDIIIIIIIIDIDIDIDDIDIIIIDDIIIIDDDDIIIDDIIDDIDIIIDDDDDDIIDIDDIIDDDDIIDDDIDIDDDIIIDIDIDIIIDDIDDDDDDDDIIDDIDDDIDDDIDDDDIIDIIIDDIDDDIDDIDDDIIDDIIIDIDIIDIDI"))

}
/*
问题： S，长度为n，只有DI，有0。。。n个数字，D代表降序，I代表升序。问有多少种合理方案

如何拆分子问题？
这个题目想法更重要些。如何拆分？将最小值的位置固定，拆出来。
我们假设最大的元素【为什么不是最小？最大的可以保证字符集不变】， 将这个最大值可能出现的所有地方的情况都加起来，就是最后结果。
这个原则就让我们拆分了问题。
假设solve(S) 就是最终的结果
如果第一个字符是D, 则最大值可以出现在第一个位置。则 拆分出的子问题是 solve(S[1:]，
【这个跟字符集剩余的字符没关系。如果有关，可以假设是最大字符，剩余的还是0 - n-1不变】

如果第一个字符是I，则不可能在第一个位置了

看中间部分， 如果是个ID 组合,比如位置i，则可以将最大值插入到这个地方。子问题不受影响。
被拆分成两个部分【去掉ID子集】，solve(S[:i-1]) * solve(S[i+1:])*可能性,有多少种可能？Cni， 组合问题，从n 选i个放前边

看最后位置，如果最后一个字符是I，则最大值可以在最后，子问题 solve(S[:len(S)-1])

如果最后一个字符是D，则不可能在最后位置了。

总的过程就是看最大值能够放在哪里？然后将问题拆分成子问题
*/
func numPermsDISequence(S string) int {
    c := nCK(len(S))
    return solve(S,make(map[string]int),c)
}
//带缓存的迭代思路
func solve(S string,cache map[string]int,nCK [][]int) int {
    mod := int(1e9)+7
    if S == "" {return 1}
    if _,ok := cache[S];ok {
        return cache[S]
    }
    ret := 0
    if S[0]=='D' { // 最大值是否能放到第一个位置
        ret += solve(S[1:],cache,nCK)
        ret %=mod
    }
    if S[len(S)-1]=='I' { //  最大值是否能放到最后位置
        ret += solve(S[:(len(S)-1)],cache,nCK)
        ret %=mod
    }
    for i:=1;i<len(S);i++ {
        // 最大值在中间部分 ID组合,处理子集去掉ID组合
        if S[i-1]=='I' && S[i]=='D' {
            left := solve(S[:(i-1)],cache,nCK)
            right := solve(S[(i+1):],cache,nCK)
            // 加法,乘法都会越界
            ret += ((left*right)%mod)*(nCK[len(S)][i]%mod)
            ret %=mod
        }
    }
    cache[S]=ret%mod
    return cache[S]
}
func nCK(n int) [][]int {
    mod := int(1e9)+7
    if n == 0 {return nil}
    // nCK = (n-1)CK + (n-1)C(K-1)
    c := make([][]int,0)
    for i:=0;i<=n;i++ {
        c = append(c,make([]int,n+1))
    }
    for i:=1;i<=n;i++ {
        c[i][0]=1
        for j:=1;j<i;j++ {
            c[i][j] = (c[i-1][j]+c[i-1][j-1])%mod
        }
        c[i][i]=1
    }
    return c
}