package main

import "fmt"

/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.
Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True
Note:
The string size will be in the range [1, 100].
*/

func main() {

    fmt.Println(checkValidString("(()"))
    fmt.Println(checkValidString("(()****"))
}
/*
暴力的方式是穷举*的三种情况，分别验证是否过得去。这个会超时。
再想个方式。其实我们关心的是配对的情况，从前往后如果把*全都当成（，看是否足够处理）；同样的，从后往前，把*全部当成）,看是否足够处理（。
其实*可空，可左，可右。有点贪心的意思，可以左右相抵的时候，*就是空,不行的时候拿*来抵挡

*/
func checkValidString(S string) bool {
    left := 0
    for i:=0;i<len(S);i++ {
        if S[i]=='('||S[i]=='*' {
            left += 1
        } else {
            left -= 1
        }
        if left < 0 {return false}
    }
    right := 0
    for i:=len(S)-1;i>=0;i-- {
        if S[i]==')'||S[i]=='*' {
            right += 1
        } else {
            right -= 1
        }
        if right < 0 {return false}
    }
    return true
}
