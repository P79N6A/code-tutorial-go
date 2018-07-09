package main

import "fmt"
/*
Given a non-negative integer n, count all numbers with unique digits, x, where 0 ≤ x < 10n.

Example:
Given n = 2, return 91. (The answer should be the total numbers in the range of 0 ≤ x < 100, excluding [11,22,33,44,55,66,77,88,99])

Credits:
*/

func countNumbersWithUniqueDigits(n int) int {
        if n == 0 {return 1}
        ret := 0
        if n > 11 {
                n = 11
        }
        for i:=1;i<=n;i++ {
                pro := 10
                if i > 1 {
                        pro = 9
                        for ii := 0; ii < i-1; ii++ {
                                pro *= (9-ii)
                        }
                }
                ret += pro
        }
        return ret
}

func main() {
        fmt.Println(countNumbersWithUniqueDigits(2))
        fmt.Println(countNumbersWithUniqueDigits(0))
}
