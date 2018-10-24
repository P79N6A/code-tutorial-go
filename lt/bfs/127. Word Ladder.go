package main

import "fmt"

/*
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/

func main() {
    fmt.Println(ladderLength("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
    
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
    exist := false
    for _,w := range wordList {
        if w == endWord {
            exist=true
            break
        }
    }
    if !exist{return 0}
    return bfs(beginWord,endWord,wordList)
}
func bfs(begin string,end string,wordlist []string) int {
    dis := 1
    queue := make([]string,0)
    queue = append(queue,begin)
    visit := make(map[string]bool)
    for len(queue)>0 {
        level := len(queue)
        for i:=0;i<level;i++ {
            if visit[queue[i]] {continue}
            if queue[i]==end {
                return dis
            }
            visit[queue[i]]=true
            for _,v := range next(queue[i],wordlist) {
                queue = append(queue,v)
            }
        }
        dis += 1
        queue = queue[level:]
    }
    return 0
}
func next(a string,wordlist []string) []string{
    ans := make([]string,0,len(wordlist))
    for _,w := range wordlist {
        if connect(a,w) {
            ans = append(ans,w)
        }
    }
    return ans
}
func connect(a,b string) bool {
    if len(a) != len(b) {return false}
    diff := 0
    for i:=0;i<len(a);i++ {
        if a[i]!=b[i]{diff+=1}
    }
    return diff==1
}