package main

import "fmt"

/*
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.

Example 1:

Input: s1 = "great", s2 = "rgeat"
Output: true
Example 2:

Input: s1 = "abcde", s2 = "caebd"
Output: false
*/

func main() {
    fmt.Println(isScramble("great","rgeat"))
    fmt.Println(isScramble("abcde","caebd"))
    fmt.Println(isScramble("abc","acb"))
}
func isScramble(s1 string, s2 string) bool {
    if len(s1) != len(s2) {return false}
    if canSame(s1,s2) == false {return false}
    if s1 == s2 {return true}
    if len(s1)<=1 {return true}
    for i:=1;i<len(s1);i++ {
        // 不交换左右子树
        if isScramble(s1[:i],s2[:i]) && isScramble(s1[i:],s2[i:]) {
            return true
        }
        // 交换左右子树
        if isScramble(s1[:i],s2[(len(s1)-i):]) && isScramble(s1[i:],s2[:(len(s1)-i)]) {
            return true
        }
    }
    return false
}
func canSame(a,b string) bool { // 是否出现了相同个数的字符，字符hash
    d := make([]int,26)
    for i:=0;i<len(a);i++ {
        d[a[i]-'a']+=1
    }
    for i:=0;i<len(b);i++ {
        d[b[i]-'a']-=1
    }
    for i:=0;i<26;i++ {
        if d[i] != 0 {return false}
    }
    return true
}