package main

import "fmt"

/*
Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/

func main() {
    //fmt.Println(findLadders("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
    fmt.Println(findLadders("qa","sq",[]string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}))
}
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    exist := false
    for _, w := range wordList {
        if w == endWord {
            exist = true
            break
        }
    }
    if !exist {
        return nil
    }
    // build graph use bfs.
    visit := make(map[string]bool)
    graph := make(map[string][]string)
    queue := make([]string,0)
    queue = append(queue,beginWord)
    visit[beginWord]=true
    for len(queue)>0 {
        t := queue[0]
        queue = queue[1:]
        graph[t]=next(t,wordList)
        for _,v := range graph[t] {
            if visit[v] == false {
                queue = append(queue,v)
                visit[v]=true
            }
        }
    }
    for k,v := range graph {
        fmt.Println(k,v)
    }

    visit = make(map[string]bool)
    ans := make([][]string, 0)
    ret := make([]string, 0)
    ret = append(ret,beginWord)
    visit[beginWord]=true
    dfs(beginWord, endWord, graph, visit, &ret, &ans)
    return ans
}
func dfs(begin, end string, graph map[string][]string, visit map[string]bool, ret *[]string, ans *[][]string) {
    if begin == end {
        r := make([]string, len(*ret))
        copy(r, *ret)
        if len(*ans)<=0 {
            *ans = append(*ans, r)
        } else {
            if len(*ret)<len((*ans)[0]) {
                *ans = [][]string{r}
            } else if len(*ret)==len((*ans)[0]) {
                *ans = append(*ans, r)
            }
        }
        return
    }
    for _,n := range graph[begin] {
        if visit[n]{continue}
        visit[n]=true
        *ret = append(*ret,n)
        dfs(n,end,graph,visit,ret,ans)
        visit[n]=false
        *ret = (*ret)[:(len(*ret)-1)]
    }
}

func next(a string, wordlist []string) []string {
    ans := make([]string, 0, len(wordlist))
    for _, w := range wordlist {
        if connect(a, w) {
            ans = append(ans, w)
        }
    }
    return ans
}
func connect(a, b string) bool {
    if len(a) != len(b) {
        return false
    }
    diff := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            diff += 1
        }
    }
    return diff == 1
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