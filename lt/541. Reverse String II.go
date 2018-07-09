package main

import "fmt"

/*
Given a string and an integer k,
you need to reverse the first k characters for every 2k characters counting
from the start of the string. If there are less than k characters left,
reverse all of them.
If there are less than 2k but greater than or equal to k characters,
then reverse the first k characters and left the other as original.
Example:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]

*/
func reverseStr(s string, k int) string {
    ret,i := "",0
    for ;i+2*k<=len(s);i += 2*k {
        ret += reverse(s,i,i+k)
        ret += s[i+k:i+2*k]
    }
    fmt.Println(i)
    if len(s) - i < k {
        ret += reverse(s,i,len(s))
    } else if len(s)-i < 2*k {
        ret += reverse(s,i,i+k)
        ret += s[i+k:]
    }
    return ret
}
func reverse(s string,i,j int) string {
    ret := make([]byte,0)
    for x:=j-1;x >=i;x--{
        ret = append(ret,s[x])
    }
    return string(ret)
}
func main() {
    fmt.Println(reverseStr("abcd",2))
    
}
