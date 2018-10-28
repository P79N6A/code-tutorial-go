package main

import (
    "strings"
    "fmt"
)

func main() {
    fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}))
}
func numUniqueEmails(emails []string) int {
    ans := make(map[string]bool)
    for _,em := range emails {
        fs := strings.Split(em,"@")
        fs[0]=strings.Replace(fs[0],".","",-1)
        idx := 0
        for ;idx<len(fs[0]);idx++ {
            if fs[0][idx] == '+' {
                break
            }
        }
        fs[0]=fs[0][:idx]
        ans[fs[0]+"@"+fs[1]]=true
    }
    return len(ans)
}