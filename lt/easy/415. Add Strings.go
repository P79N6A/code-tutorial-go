package main

import "fmt"

/*

Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

*/
func addStrings(num1 string, num2 string) string {
    var carry byte
    if len(num2)>len(num1) {
        num2,num1 = num1,num2
        }
    ret := ""
    l1,l2 := len(num1),len(num2)
    for i:=1;i<=l1;i++ {
        sum := num1[l1-i]-'0' + carry
        if i <= l2 {
            sum = num1[l1-i]-'0' + num2[l2-i]-'0' + carry
        }
        carry = sum/10
        ret = string(sum%10+'0')+ret
    }
    if carry > 0 {
        ret = "1"+ret
    }
    return ret
}
func main() {
    //fmt.Println(addStrings("123","888"))
    fmt.Println(addStrings("1","888"))
    fmt.Println(addStrings("222222","888"))

}
