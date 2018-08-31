package main

import (
    "math"
    "fmt"
)

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
*/

func main() {
    fmt.Println(divide(math.MinInt32,-1))
}
func divide(dividend int, divisor int) int {
    /*
    不能用乘法除法和取模。
    直接想到的方式是减法，但收敛太慢；线性的。
    减法配合<<1 位运算，这样做减法做的更快，看能够被减掉多少次。
    */
    if dividend == math.MinInt32 && divisor == -1 {
        // 就这么个特殊情况。
        return math.MaxInt32
    }
    if dividend>0&&divisor<0 {
        return -1*solve(dividend,-1*divisor)
    } else if divisor>0&&dividend<0 {
        return -1*solve(-1*dividend,divisor)
    } else if dividend<0&&divisor<0 {
        return solve(-1*dividend,-1*divisor)
    }
    return solve(dividend,divisor)}
/*
Well, two cases may cause overflow:

divisor = 0;
dividend = INT_MIN and divisor = -1 (because abs(INT_MIN) = INT_MAX + 1).
*/

func solve(dd,ds int) int {
    if dd<ds {
        return 0
    }
    var x int = 1
    var t int = ds
    for dd > t<<1 {
        t<<=1
        x<<=1  // 2倍关系，logN的复杂度
    }
    //fmt.Println(x,t,dd,ds)
    return x + solve(dd-t,ds)
}