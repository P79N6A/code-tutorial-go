package main

import (
    "strconv"
    "fmt"
)

/*
Given a string containing only digits,
restore it by returning all possible valid IP address combinations.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/
func restoreIpAddresses(s string) []string {
    return restoreIP(0,s)
}
func restoreIP(k int,s string) []string {
    if k == 3 {
        // 最后一位
        i,e := strconv.Atoi(s)
        if e != nil || (len(s)>1 &&s[0]=='0') {return nil}
        if i <= 255 {
            return []string{s}
        }
    }
    ret := make([]string,0)
    for i:=1;i<=3&&i<len(s);i++ {
        if i > 1 && s[0]=='0'{continue}
        ii,e := strconv.Atoi(s[:i])
        if e != nil || ii > 255 {continue}
        rr := restoreIP(k+1,s[i:])
        for _,r := range rr {
            ret = append(ret,s[:i] + "." + r)
        }
    }
    return ret
}
func main() {
    //fmt.Println(restoreIpAddresses("25525511135"))
    //fmt.Println(restoreIpAddresses("010010"))
    fmt.Println(restoreIpAddresses("255255255255"))

}
