package main

import "sort"

/*
We are given two arrays A and B of words.  Each word is a string of lowercase letters.

Now, say that word b is a subset of word a if every letter in b occurs in a, including multiplicity.  For example, "wrr" is a subset of "warrior", but is not a subset of "world".

Now say a word a from A is universal if for every b in B, b is a subset of a.

Return a list of all universal words in A.  You can return the words in any order.



Example 1:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
Output: ["facebook","google","leetcode"]
Example 2:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
Output: ["apple","google","leetcode"]
Example 3:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
Output: ["facebook","google"]
Example 4:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
Output: ["google","leetcode"]
Example 5:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
Output: ["facebook","leetcode"]


Note:

1 <= A.length, B.length <= 10000
1 <= A[i].length, B[i].length <= 10
A[i] and B[i] consist only of lowercase letters.
All words in A[i] are unique: there isn't i != j with A[i] == A[j].
*/

func main() {
    
}
/*
这道题目逻辑相对简单，就是看两个字符串是否是包含关系；转换成map比较即可
问题难在A,B都是两个集合。题目说A是唯一的。如何避免重复计算？
第一种方法 ：就是把B中"相同"的字符合并成一个。 这个相同指的是字符出现的次数一致。
第二种方法：既然要满足B中全部的，则取各个字符出现的最大值即可。！！！！！！
*/
func wordSubsets(A []string, B []string) []string {
    maxB := make([]int,26)
    for _,b := range B {
        bd := make([]int,26)
        for i:=0;i<len(b);i++ {
            bd[b[i]-'a']+=1
        }
        for i:=0;i<26;i++ {
            if maxB[i]<bd[i] {
                maxB[i]=bd[i]
            }
        }
    }
    ans := make([]string,0)
    outer:
    for _,a := range A {
        ad := make([]int,26)
        for i:=0;i<len(a);i++ {
            ad[a[i]-'a']+=1
        }
        for i:=0;i<26;i++ {
            if ad[i]<maxB[i] {
                continue outer
            }
        }
        ans = append(ans,a)
    }
    return ans
}
func wordSubsetsi(A []string, B []string) []string {
    adict := make(map[string]map[int]int)
    for _,AA := range A {
        dict := make(map[int]int)
        for i:=0;i<len(AA);i++ {
            dict[int(AA[i]-'a')] += 1
        }
        adict[AA]=dict
    }
    bdict := make(map[string]map[int]int)
    BBB := make(map[string]bool)
    for _,b := range B {
        // 最关键的地方是对B做归一化，否则会超时。
        // 其实还是在如何判断字符串相等。。
        bs := []byte(b)
        sort.Slice(bs, func(i, j int) bool {
            return bs[i]<bs[j]
        })
        BBB[string(bs)]=true
    }
    for BB,_ := range BBB {
        dict := make(map[int]int)
        for i:=0;i<len(BB);i++ {
            dict[int(BB[i]-'a')] += 1
        }
        bdict[BB]=dict
    }
    ans := make([]string,0)
    for k,a := range adict {
        issub := true
        for _,b := range bdict {
            if subsetDict(a,b) == false {
                issub = false
                break
            }
        }
        if issub {
            ans = append(ans,k)
        }
    }
    return ans
}
func subsetDict(A,B map[int]int) bool {
    for k,v := range B {
        if A[k] < v {
            return false
        }
    }
    return true
}
func subset(A,B string) bool {
    dict := make([]int,26)
    for i:=0;i<len(A);i++ {
        dict[A[i]-'a'] += 1
    }
    for i:=0;i<len(B);i++ {
        dict[B[i]-'a'] -= 1
        if dict[B[i]-'a'] < 0 {
            return false
        }
    }
    return true
}