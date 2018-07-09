package main

import "fmt"

/*

Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/
func myPow(x float64, n int) float64 {
    if n == 1 {return x}
    if n == 0 {return 1}
    if n < 0 {
        x = 1/x
        n = -n
    }
    if n%2 == 0 {
        return myPow(x*x,n/2)
    }
    return x*myPow(x*x,n/2)
}
func main() {
    fmt.Println(myPow(2.1,3))
    fmt.Println(myPow(2.0,10))

}