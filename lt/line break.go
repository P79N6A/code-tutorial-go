package main

import (
    "fmt"
    "strings"
)

// dp[i]=max(dp[j] + c(j-i) one line)
func lineBeaker2(senten string, row int) int {

}
func lineBeaker(senten string, row int) int {
    words := strings.Split(senten," ")
    r,l := row,0
    rets := make([]string,0)
    line := ""
    for _,w := range words {
        if r < len(w) {
            rets = append(rets,line)
            l,r,line = l+1,row,""
        }
        r -= len(w) + 1
        line += w + " "
    }
    if r > 0 {
        rets = append(rets,line)
        l += 1
    }
    for _,r := range rets {
        fmt.Println(r)
    }
    return l
}

func main() {
    fmt.Println(lineBeaker("I'm a good guy, and I know what I should not to do!",25))
    
}
