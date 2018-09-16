package main

import (
        "strconv"
        "fmt"
)

func main() {
        fmt.Println(genPalindrome(3))
        fmt.Println(superpalindromesInRange("4", "1000"))
}
func superpalindromesInRange(L string, R string) int {
        l64, _ := strconv.ParseInt(L, 10, 0)
        r64, _ := strconv.ParseInt(R, 10, 0)
        ret := 0
        var i int64
        for ii := 1; ii < 20; ii++ {
                for _, i = range genPalindrome(ii) {
                        if i * i > r64 {
                                goto end
                        }
                        if i * i < l64 {
                                continue
                        }

                        if ispalin(strconv.FormatInt(i * i, 10)) {
                                fmt.Println(i * i)
                                ret += 1
                        }
                }
        }
        end:
        return ret
}
func ispalin(r string) bool {
        if len(r) <= 1 {
                return true
        }
        if r[0] == r[len(r) - 1] {
                return ispalin(r[1:(len(r) - 1)])
        }
        return false
}

func genPalindrome(n int) []int64 {
        // TODO.
        if n == 0 {
                return nil
        }
        if n == 1 {
                return []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
        }
        if n == 2 {
                return []int64{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}
        }
        ret := make([]int64, 0)
        sub := genPalindrome(n - 2)
        //ret = append(ret,0)
        var i int64
        for i = 0; i <= 9; i++ {
                for _, s := range sub {
                        r := i
                        for j := 1; j < n; j++ {
                                r *= 10
                        }
                        rt := r + s * 10 + i
                        ret = append(ret, rt)
                }
        }
        return ret
}
