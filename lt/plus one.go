package main

import "fmt"

func plusOne (digits []int) []int {
    // write your code here
    carry := 1
    for i:=len(digits)-1;i>=0;i-- {
        s := (digits[i]+carry)%10
        carry = (digits[i]+carry)/10
        digits[i]=s
    }
    if carry > 0 {
        digits = append(digits,0)
        copy(digits[1:],digits[:len(digits)-1])
        digits[0]=carry
    }
    return digits
}


func main() {
    fmt.Println(plusOne([]int{9}))
    
}
