package main

import "fmt"

/*
Given a string which contains only lowercase letters, remove duplicate letters so that every letter appear once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.

Example 1:

Input: "bcabc"
Output: "abc"
Example 2:

Input: "cbacdcbc"
Output: "acdb"

*/
func removeDuplicateLetters(s string) string {
    freq := make(map[byte]int)
    for i:=0;i<len(s);i++ {freq[s[i]]+=1}
    visit := make(map[byte]bool)
    res := make([]byte,0)
    for i:=0;i<len(s);i++ {
        // 记录词频
        freq[s[i]] -=1
        if visit[s[i]] {continue}
        //一般都是先处理出栈，入栈是一定要做的，放到后边
        for len(res)>0 && s[i] < res[len(res)-1] && freq[res[len(res)-1]]>0 {
            // 如果栈不空，并且当前字符小于栈顶，并且栈顶元素在后边还有，则出栈，并标记未访问
            visit[res[len(res)-1]]=false //先标记，下边再删掉
            res = res[:len(res)-1]
        }
        //如果当前字符大于栈顶，入栈
        res = append(res,s[i])
        visit[s[i]]=true
    }
    return string(res)
}
func main() {
    fmt.Println(removeDuplicateLetters("cbbbcaa"))
    fmt.Println(removeDuplicateLetters("bbcaac"))
    fmt.Println(removeDuplicateLetters("bcabc"))
    fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
