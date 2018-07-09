package main

import "fmt"

/*

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
Example 3:

Input: "IX"
Output: 9
Example 4:

Input: "LVIII"
Output: 58
Explanation: C = 100, L = 50, XXX = 30 and III = 3.
Example 5:

Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
func romanToInt(s string) int {
    ret := 0
    var next byte
    for i:=0;i<len(s);i++ {
        if i < len(s)-1 {
            next=s[i+1]
        } else {
            next = 'a'
        }
        switch s[i] {
        case 'I':
            i+=1
            if next == 'V' {
                ret += 4
            } else if next == 'X' {
                ret += 9
            } else {
                ret += 1
                i-=1
            }
        case 'X':
            i+=1
            if next == 'L' {
                ret += 40
            } else if next == 'C' {
                ret += 90
            } else {
                ret += 10
                i-=1
            }
        case 'C':
            i+=1
            if next == 'D' {
                ret += 400
            } else if next == 'M' {
                ret += 900
            } else {
                ret += 100
                i-=1
            }
        case 'M':
            ret +=1000
        case 'D':
            ret += 500
        case 'L':
            ret += 50
        case 'V':
            ret += 5
        }
    }
    return ret
}
func main() {
    fmt.Println(romanToInt("DCXXI"))
    
}
