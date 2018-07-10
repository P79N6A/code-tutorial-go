package main

import "fmt"
/*
We have a string S of lowercase letters, and an integer array shifts.

Call the shift of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').

For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.

Now for each shifts[i] = x, we want to shift the first i+1 letters of S, x times.

Return the final string after all such shifts to S are applied.

Example 1:

Input: S = "abc", shifts = [3,5,9]
Output: "rpl"
Explanation:
We start with "abc".
After shifting the first 1 letters of S by 3, we have "dbc".
After shifting the first 2 letters of S by 5, we have "igc".
After shifting the first 3 letters of S by 9, we have "rpl", the answer.
Note:

1 <= S.length = shifts.length <= 20000
0 <= shifts[i] <= 10 ^ 9

*/
func shiftingLetters(S string, shifts []int) string {
        sum := make([]int,len(shifts)+1)
        for i:=len(shifts)-1;i>=0;i-- {
                sum[i] = sum[i+1]+shifts[i]
        }
        ret := ""
        fmt.Println(sum)
        for i:=0;i<len(S);i++ {
                n := S[i] + byte(sum[i]%26)
                if n > 'z' {
                        n = 'a'+n-'z'-1
                }
                ret += string(n)
        }
        return ret
}

func main() {
        //fmt.Println(shiftingLetters("abz",[]int{3,5,1}))
        //fmt.Println('z'-'a')
        fmt.Println((411)%26)
        fmt.Println(string('x'+21-26))

        //"surjgj"
        fmt.Println(shiftingLetters("xrdofd", []int{70,41,71,72,73,84}))
}
