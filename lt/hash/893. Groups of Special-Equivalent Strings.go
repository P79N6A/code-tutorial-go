package main

import "fmt"

/*
You are given an array A of strings.

Two strings S and T are special-equivalent if after any number of moves, S == T.

A move consists of choosing two indices i and j with i % 2 == j % 2, and swapping S[i] with S[j].

Now, a group of special-equivalent strings from A is a non-empty subset S of A such that any string not in S is not special-equivalent with any string in S.

Return the number of groups of special-equivalent strings from A.



Example 1:

Input: ["a","b","c","a","c","c"]
Output: 3
Explanation: 3 groups ["a","a"], ["b"], ["c","c","c"]
Example 2:

Input: ["aa","bb","ab","ba"]
Output: 4
Explanation: 4 groups ["aa"], ["bb"], ["ab"], ["ba"]
Example 3:

Input: ["abc","acb","bac","bca","cab","cba"]
Output: 3
Explanation: 3 groups ["abc","cba"], ["acb","bca"], ["bac","cab"]
Example 4:

Input: ["abcd","cdab","adcb","cbad"]
Output: 1
Explanation: 1 group ["abcd","cdab","adcb","cbad"]


Note:

1 <= A.length <= 1000
1 <= A[i].length <= 20
All A[i] have the same length.
All A[i] consist of only lowercase letters.
*/

func main() {
    fmt.Println(numSpecialEquivGroups([]string{"abcd","cdab","adcb","cbad"}))
    fmt.Println(numSpecialEquivGroups2([]string{"abcd","cdab","adcb","cbad"}))

}
/*
问题：将给出的字符串数组分组，同组的字符串在奇偶位置的字符集合一致
思考：对于字符集合（有重复字符）创建hash如何才能更简练？
map[byte]int？  [26]int? 这种数组类型在比较的时候总会不方便。
利用下 []byte<=>string
*/
func numSpecialEquivGroups(A []string) int {
    // map的key不能是slice,但可以是array
    set := make(map[string]bool)
    for _,s := range A {
        cnt := make([]byte,52)// 分成两段，做记录
        for i:=0;i<len(s);i++ {
            cnt[s[i]-'a'+uint8(26*(i%2))]+=1
        }
        // 最终转成 string, 方便放到set里面去。
        set[string(cnt)]=true
    }
    return len(set)
}
func numSpecialEquivGroups2(A []string) int {
    // map的key不能是slice,但可以是array
    set2 := make(map[[52]byte]bool) // 其实这是个双map结构
    for _,s := range A {
        cnt := [52]byte{} //
        for i:=0;i<len(s);i++ {
            // 创建单个字符串的字符map
            cnt[s[i]-'a'+uint8(26*(i%2))]+=1
        }
        set2[cnt]=true
    }
    return len(set2)
}
