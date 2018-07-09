package main

import (
    "strconv"
    "fmt"
)

/*

Write a function to generate the generalized abbreviations of a word.

Example:

Given word = "word", return the following list (order does not matter):

["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]
*/
func generateAbbreviations(word string) []string {
    // return solve(word,0,len(word))
    ret := make([]string,0)
    dfs(word,"",&ret)
    return ret
}
func dfs(word string, str string, ret *[]string) {
    if len(word)==0 {
        *ret = append(*ret,str)
        return
    }
    // 第一种就是只拆出来一个字符
    dfs(word[1:],str+string(word[0]),ret)
    for i:=1;i<len(word);i++ {
        //中间情况，但要保证最后一个是字符，否则会重复
        dfs(word[i+1:],str+strconv.Itoa(i)+string(word[i]),ret)
    }
    // 最后情况就是全部变成数字
    dfs("",str+strconv.Itoa(len(word)),ret)
}
func solve2(word string, pos int,w string, ret *[]string) {
    if pos >= len(word) {
        fmt.Println(w)
        *ret = append(*ret,w)
        return
    }
    fmt.Println(w)
    for i:=pos+1;i<len(word);i++ {
        solve2(word,pos+1,w,ret)
    }
    solve2(word,len(word),w+strconv.Itoa(len(word)),ret)

}
func solve(word string, i,j int) []string {
    if i+1 == j {
        return []string{"1",string(word[i])}
    }
    ret := make([]string,0)
    for ii:=i+1;ii<j;ii++ {
        l := solve(word,i,ii)
        r := solve(word,ii,j)
        for _,ll := range l {
            for _,rr := range r {
                if ll[len(ll)-1] >= '0' && ll[len(ll)-1] <= '9' && rr[0] >= '0' && rr[0] <= '9' {
                    continue
                }
                ret = append(ret,ll+rr)
            }
        }
    }
    ret = append(ret,strconv.Itoa(j-i))
    return ret
}

func main() {
    fmt.Println(generateAbbreviations("word"))
    
}
