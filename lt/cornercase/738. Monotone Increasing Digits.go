package main

import (
    "strconv"
    "fmt"
)

/*
Given a non-negative integer N, find the largest number that is less than or equal to N with monotone increasing digits.

(Recall that an integer has monotone increasing digits if and only if each pair of adjacent digits x and y satisfy x <= y.)

Example 1:
Input: N = 10
Output: 9
Example 2:
Input: N = 1234
Output: 1234
Example 3:
Input: N = 332
Output: 299
Note: N is an integer in the range [0, 10^9].


*/

func main() {
    fmt.Println(monotoneIncreasingDigits(157392))
    fmt.Println(monotoneIncreasingDigits(3333222))
    fmt.Println(monotoneIncreasingDigits(1234))
}
func monotoneIncreasingDigits(N int) int {
    str := strconv.Itoa(N)
    up,down := 0,-1
    for i:=1;i<len(str);i++ {
        if str[i-1]<str[i] {
            up = i  // 下降趋势之前的最后一个上升趋势，因为会有相等情况，比如22211
        } else if str[i-1]>str[i] {
            down=i
            break
        }
    }
    ans := str
    if down > 0 { // 存在下降趋势
        ans = str[:up]
        ans += string(str[up]-1)
        for i:=up+1;i<len(str);i++ {
            ans += "9"
        }
    }
    a,_ := strconv.Atoi(string(ans))
    return a
}
func monotoneIncreasingDigits1(N int) int {
    if N < 10 {return N}
    str := strconv.Itoa(N)
    i,j := 0,1
    for j<len(str) {
        if str[i] < str[j] {
            i = j  //  会重复。。。
            j += 1
        } else if str[i]==str[j] {
            j+=1
        } else {
            break
        }
    }
    ans := str
    if j < len(str) {
        ans = str[:i]
        ans += string(str[i]-1)
        for i+=1;i<len(str);i++ {
            ans+="9"
        }
    }
    a,_ := strconv.Atoi(string(ans))
    return a
}