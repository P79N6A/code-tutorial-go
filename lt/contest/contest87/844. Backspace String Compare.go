package main

import "fmt"
/*
Given two strings S and T, return if they are equal when both are typed into empty text editors.
 # means a backspace character.



Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".


Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
*/
func backspaceCompare(S string, T string) bool {
        return backspace(S) == backspace(T)
}
func backspace(str string) string {
        ret := ""
        for i:=0;i<len(str);i++ {
                if str[i] == '#'  {
                        if len(ret) > 0 {
                                ret = ret[:len(ret)-1]
                        }
                } else {
                        ret += string(str[i])
                }
        }
        return ret
}
func main() {
        fmt.Println(backspaceCompare("y#fo##f","y#f#o##f"))
        fmt.Println(backspace("y#fo##f"))
        fmt.Println(backspace("y#f#o##f"))

}
