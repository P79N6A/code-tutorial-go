package main

import (
    "strconv"
    "fmt"
    "math"
)

/*
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

Example 1:

Input: numerator = 1, denominator = 2
Output: "0.5"
Example 2:

Input: numerator = 2, denominator = 1
Output: "2"
Example 3:

Input: numerator = 2, denominator = 3
Output: "0.(6)"

*/

func main() {
    /*
    fmt.Println(fractionToDecimal(1, 2))
    fmt.Println(fractionToDecimal(2, 1))
    fmt.Println(fractionToDecimal(2, 3))
    fmt.Println(fractionToDecimal(10, 3))
    fmt.Println(fractionToDecimal(10, 7))
    fmt.Println(fractionToDecimal(2, 2))
    */
    fmt.Println(fractionToDecimal(4, 3333))
    fmt.Println(fractionToDecimal(111, 1000))
}
func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {return "0"}
    ans := ""
    // +  or -
    if numerator*denominator < 0 {
        ans += "-"
    }
    // remove sig
    num := int(math.Abs(float64(numerator)))
    den := int(math.Abs(float64(denominator)))
    // interger part
    ans += strconv.Itoa(num/den)
    num %=den
    if num == 0 {return ans}
    // fractional part
    ans += "."
    dict := make(map[int]int)
    dict[num]=len(ans)
    for num != 0 {
        num *= 10
        ans += strconv.Itoa(num/den) // 不论如何，都需要将商放到结果中去
        num %= den  //
        if _,ok := dict[num];ok {
            ans = ans[:dict[num]] + "(" + ans[dict[num]:] + ")"
            break
        } else {
            dict[num] = len(ans)
        }
    }
    return ans
}
