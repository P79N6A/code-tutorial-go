package main

import (
    "fmt"
    "sort"
)

/*
Given a list accounts, each element accounts[i] is a list of strings, where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some email that is common to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails in sorted order. The accounts themselves can be returned in any order.

Example 1:
Input:
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
Output: [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
Explanation:
The first and third John's are the same person as they have the common email "johnsmith@mail.com".
The second John and Mary are different people as none of their email addresses are used by other accounts.
We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.
Note:

The length of accounts will be in the range [1, 1000].
The length of accounts[i] will be in the range [1, 10].
The length of accounts[i][j] will be in the range [1, 30].
*/

func main() {
    ans := accountsMerge([][]string{
        {"John", "johnsmith@mail.com", "john00@mail.com"},
        {"John", "johnnybravo@mail.com"},
        {"John", "johnsmith@mail.com", "john_newyork@mail.com"},
        {"Mary", "mary@mail.com"},
    })
    for _,a := range ans {
        fmt.Println(a)
    }
}

func accountsMerge(accounts [][]string) [][]string {
    graph := make(map[string][]string) // 构建邻接表
    for _,ac := range accounts {
        for i:=1;i<len(ac);i++ {
            if len(graph[ac[i]]) <= 0 {
                graph[ac[i]] = ac  // 把名字放到第一个
            } else {
                graph[ac[i]]=append(graph[ac[i]],ac[1:]...) // 去掉名字，有重复没关系
            }
        }
    }
    ret := make([][]string,0)
    visit := make(map[string]bool)
    for k,v := range graph {
        if visit[k] == false { // 发现一个未访问的连通图
            visit[k]=true
            ans := []string{v[0],k} // 保存联通图节点列表
            dfs(k,graph,visit,&ans) // dfs直接处理完一个连通图
            sort.Strings(ans)
            ret = append(ret,ans)
        }
    }
    return ret
}
func dfs(u string,graph map[string][]string,visit map[string]bool,ans *[]string) {
    for i:=1;i<len(graph[u]);i++ {
        if visit[graph[u][i]] == false {
            *ans = append(*ans,graph[u][i])
            visit[graph[u][i]] = true
            dfs(graph[u][i],graph,visit,ans)
        }
    }
}
