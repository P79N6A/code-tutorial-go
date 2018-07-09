package main

import "fmt"

func myAtoi(str string) int {
        i := 0
        var  neg int64 = 1
        for ;i<len(str);i++{
                if str[i] == ' ' {
                        continue
                } else {
                        if str[i] == '-' {
                                neg = -1
                                i += 1
                        } else if str[i] == '+' {
                                neg = 1
                                i += 1
                        }
                        break
                }
        }
        var res int64 = 0
        str = str[i:]
        fmt.Println(str)
        for i = 0;i < len(str);i++ {
                if str[i] < '0' || str[i] > '9' {
                        break
                }
                res = res * 10 + int64(str[i]-'0')
                if res > 2147483647 || res < -2147483648 {
                        break
                }
        }
        fmt.Println(res,neg)
        res = neg * res
        if res > 2147483647 {return 2147483647}
        if res < -2147483648 {return -2147483648}
        return int(res)
}





func main() {
        fmt.Println(myAtoi("9223372036854775809"))
}
