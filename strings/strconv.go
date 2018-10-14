package main

import (
        "strconv"
        "fmt"
)

func main() {
        fmt.Println(strconv.FormatInt(700,26))
        fmt.Println(trailingZeroes(30))
}
func trailingZeroes(n int) int {
        ans := 1
        zero := 0
        for n>1 {
                ans *= n
                n-=1
                for ans%10 == 0 {
                        ans /= 10
                        zero += 1
                }
        }
        for ans%10 == 0 {
                ans /= 10
                zero += 1
        }
        return zero
}