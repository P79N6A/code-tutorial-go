package main

import "fmt"

/*

*/
func canWin (s string) bool {
    // write your code here
    for i:=0;i+1<len(s);i++ {
        if s[i]=='+' && s[i+1]=='+' {
            newS:=s[:i]+"--"+s[i+2:]
            if !canWin(newS) {
                return true
            }
        }
    }
    fmt.Println(s)
    return false
}
func main() {
    fmt.Println(canWin("+++++++++"))
}
