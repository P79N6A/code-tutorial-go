package main

import (
    "fmt"
    "strconv"
    "sort"
)
/*
You have an initial power P, an initial score of 0 points, and a bag of tokens.

Each token can be used at most once, has a value token[i], and has potentially two ways to use it.

If we have at least token[i] power, we may play the token face up, losing token[i] power, and gaining 1 point.
If we have at least 1 point, we may play the token face down, gaining token[i] power, and losing 1 point.
Return the largest number of points we can have after playing any number of tokens.



Example 1:

Input: tokens = [100], P = 50
Output: 0
Example 2:

Input: tokens = [100,200], P = 150
Output: 1
Example 3:

Input: tokens = [100,200,300,400], P = 200
Output: 2


Note:

tokens.length <= 1000
0 <= tokens[i] < 10000
0 <= P < 10000
*/
func main() {
    fmt.Println(bagOfTokensScore([]int{100},50))//0
    fmt.Println(bagOfTokensScore([]int{100,200},150))//1
    fmt.Println(bagOfTokensScore([]int{100,200,300,400},200))//2
    fmt.Println(bagOfTokensScore([]int{52,65,35,88,28,1,4,68,56,95},94))
    fmt.Println(bagOfTokensScore([]int{87,24,32},87))//1
    //*/
    //fmt.Println(bagOfTokensScore([]int{48,26,87},81))//2
    //fmt.Println(bagOfTokensScore([]int{35,18,50,3},10))//2
    fmt.Println(bagOfTokensScore([]int{52,65,35,88,28,1,4,68,56,95},94))//5

}
/*
写了几个解法,比如dfs穷举;最终确定就是贪心解法;因为可以用小的power换取相同的point
问题:就是用手里的Power换point.
相同的point可以还更大的power.
原则:贪心加双指针;用power还小的token得到point;用point得到大的power
将手里的power都兑换完，看能得到多少分。如果当前分>0，就换得最大的power继续这个过程。在换取完毕，记录下max
*/
func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    ans := 0
    points := 0
    left,right := 0,len(tokens)-1
    for left <= right && (power>=tokens[left] ||points>0) {
        for left <= right&&power>=tokens[left]{
            // 从小到大,left++ 可以得到point,衰减power
            power-=tokens[left]
            left += 1
            points += 1
        }
        ans = max(ans,points)
        if left<=right && points > 0{
            // 从大到小,right--,可以增大power,衰减point
            power += tokens[right]
            right -= 1
            points -= 1
        }
    }
    return ans
}
func bagd(tokens []int,P int,point int) int {
    fmt.Println(P,point,tokens)
    if len(tokens) <= 0 {return point}
    if P < tokens[0] {return point}
    if P >=tokens[len(tokens)-1] {
        return bagd(tokens[1:],P-tokens[0],point+1)
    }
    idx := bslessOne(tokens,P)
    nt := make([]int,0)
    nt = append(nt,tokens[:idx]...)
    nt = append(nt,tokens[(idx+1):(len(tokens)-1)]...)
    return bagd(nt,tokens[len(tokens)-1]+P-tokens[idx],point)
}
func bagOfTokensScore2(tokens []int, P int) int {
    sort.Ints(tokens)
    x := 0
    pp := P
    for i:=0;i<len(tokens);i++ {
        pp -= tokens[i]
        if pp >= 0 {x += 1}
    }
    fmt.Println(tokens,x)
    np,nt := maxP(tokens,P)
    fmt.Println(np,nt)
    ans := 0
    for i:=0;i<len(nt);i++ {
        np -= nt[i]
        if np >= 0 {ans += 1}
    }
    fmt.Println(ans)
    return max(ans,x)
}
func bslessOne(data []int,key int) int {
    l,r := 0,len(data)
    for l < r {
        m := (l+r)/2
        if data[m] == key {
            return m
        } else if data[m] < key {
            l = m+1
        } else {
            r = m
        }
    }
    return r-1
}
func maxP(tokens []int,P int) (int,[]int) {
    if len(tokens)<=0 {return P,tokens}
    if P > tokens[len(tokens)-1] {
        return P,tokens
    }
    if P < tokens[0] {return P,tokens}
    idx := bslessOne(tokens,P)
    nt := make([]int,0)
    nt = append(nt,tokens[:idx]...)
    if idx < len(tokens)-1 {
        nt = append(nt,tokens[(idx+1):(len(tokens)-1)]...)
    }
    if len(nt) <= 0 {
        return P,tokens
    }
    return tokens[len(tokens)-1]+P-tokens[idx],nt
}
func bagOfTokensScore1(tokens []int, P int) int {
    ans := 0
    visit := make([]byte,len(tokens))
    for i:=0;i<len(visit);i++ {visit[i]='0'}
    solve(tokens,&visit,P,0,make(map[string]int),&ans)
    return ans
}
func solve(tokens []int,visit *[]byte,power,point int,cache map[string]int, ans *int) {
    key := string(*visit) + "-" + strconv.Itoa(power) + "-" + strconv.Itoa(point)
    fmt.Println(key)
    if _,ok := cache[key];ok {return}
    *ans = max(*ans,point)
    for i:=0;i<len(tokens);i++ {
        if (*visit)[i]=='1' {continue}
        (*visit)[i]='1'
        solve(tokens,visit,power,point,cache,ans)
        if power >= tokens[i] {
            solve(tokens,visit,power-tokens[i],point+1,cache,ans)
        }
        if point > 0 {
            solve(tokens,visit,power+tokens[i],point-1,cache,ans)
        }
        (*visit)[i]='0'
    }
}
func max(a,b int) int {
    if a>b {return a}
    return b
}