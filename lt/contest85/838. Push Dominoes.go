package main

import (
        "fmt"
        "strings"
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
        dlen := len(dominoes)
        result,i,j := "",0,0
        for i < dlen {
                if dominoes[i] == '.' {
                        i += 1
                }
        }
        if dominoes[i] == 'L' {
                result += strings.Repeat("L",i)
        }
        for j < dlen && i < dlen {
                for i < dlen && dominoes[i] != 'R' {
                        i += 1
                        result += string(dominoes[i])
                }
                j = i
                for j < dlen && dominoes[j] != 'L' {
                        j += 1
                }
                // set i-j domino

        }



        return dominoes
}

func main() {
        fmt.Println(pushDominoes(".RL"))
}
