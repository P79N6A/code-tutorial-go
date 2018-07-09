package main

import (
    "strconv"
    "strings"
    "fmt"
)

/*

Additive number is a string whose digits can form additive sequence.

A valid additive sequence should contain at least three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

Given a string containing only digits '0'-'9', write a function to determine if it's an additive number.

Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.

Example 1:

Input: "112358"
Output: true
Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5, 8.
             1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
Example 2:

Input: "199100199"
Output: true
Explanation: The additive sequence is: 1, 99, 100, 199.
             1 + 99 = 100, 99 + 100 = 199
Follow up:
How would you handle overflow for very large input integers?
*/
func isAdditiveNumber(num string) bool {
    for i:=0;i<len(num);i++ {
        // 判断是否有前置0
        if i > 0 && num[0]=='0' {
            continue
        }
        one,_ := strconv.Atoi(num[:i+1])
        for j:=i+2;j<len(num);j++ {
            // 判断是否有前置0
            if j-i-1 > 1 && num[i+1]=='0' {
                continue
            }
            two,_ := strconv.Atoi(num[i+1:j])
            if bt(one,two,num[j:]) {
                return true
            }
        }
    }
    return false
}
func bt(one,two int, num string) bool {
    if len(num) <= 0 {
        return true
    }
    f := strconv.Itoa(one + two)
    if strings.HasPrefix(num,f) {
        return bt(two,one+two,num[len(f):])
    }
    return false
}

func main() {
    fmt.Println(isAdditiveNumber("199100199"))
    fmt.Println(isAdditiveNumber("112358"))
    fmt.Println(isAdditiveNumber("000"))
    fmt.Println(isAdditiveNumber("0"))
    fmt.Println(isAdditiveNumber("00"))
    fmt.Println(isAdditiveNumber("012"))
    fmt.Println(isAdditiveNumber("0235813"))
}
