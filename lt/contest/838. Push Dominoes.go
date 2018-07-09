package main

import (
    "strings"
    "fmt"
)

/*
There are N dominoes in a line, and we place each domino vertically upright.

In the beginning, we simultaneously push some of the dominoes either to the left or to the right.



After each second, each domino that is falling to the left pushes the adjacent domino on the left.

Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

Given a string "S" representing the initial state. S[i] = 'L', if the i-th domino has been pushed to the left; S[i] = 'R', if the i-th domino has been pushed to the right; S[i] = '.', if the i-th domino has not been pushed.

Return a string representing the final state.

Example 1:

Input: ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."
Example 2:

Input: "RR.L"
Output: "RR.L"
Explanation: The first domino expends no additional force on the second domino.
Note:

0 <= N <= 10^5
String dominoes contains only 'L', 'R' and '.'
*/

func pushDominoes(dominoes string) string {
    l,r := 0,0
    ret := ""
    for r < len(dominoes) {
        if dominoes[r] == '.' {
            r += 1
        } else {
            if dominoes[l]=='R' && dominoes[r]=='L' {
                // R...L
                ret += strings.Repeat("R",(r-l+1)/2)
                if (r-l)%2 == 0 {
                    ret += "."
                    ret += strings.Repeat("L",(r-l)/2-1)
                } else {
                    ret += strings.Repeat("L",(r-l)/2)
                }
            } else if dominoes[r] == 'R' && dominoes[l]=='R' {
                // R..R
                ret += strings.Repeat("R",r-l)
            } else  if dominoes[r]=='L' && dominoes[l]=='L' {
                // L...L
                ret += strings.Repeat("L",r-l)
            } else if dominoes[l] =='L' && dominoes[r]=='R' {
                // L...R
                ret += dominoes[l:r]
            } else if dominoes[r]=='L' && l <= 0 {
                l = 0
                // ....L
                ret += strings.Repeat("L",r-l)
            } else if dominoes[r] == 'R' && l <= 0 {
                l = 0
                // ....R
                ret += dominoes[:r]
            }
            l,r = r,r+1
        }
    }
    if dominoes[l]=='R' {
        // R.....
        ret += strings.Repeat("R",r-l)
    } else {
        // L..... or other
        ret += dominoes[l:]
    }
    return ret
}
func main() {
    //fmt.Println(pushDominoes(".L.R...LR..L.."))
    fmt.Println(pushDominoes("R.R.L"))
    //fmt.Println(pushDominoes(".L.R."))
    //fmt.Println(pushDominoes("."))
    //fmt.Println(pushDominoes("RL"))
}
