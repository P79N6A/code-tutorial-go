package main

import "fmt"

/*

Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

For example, given n = 2, return 1 (2 = 1 + 1); given n = 10, return 36 (10 = 3 + 3 + 4).

Note: You may assume that n is not less than 2 and not larger than 58.
*/
/*
如何去思考这个问题？先看递归，比如说f(n)是乘积最大的，那么
f(n)=max(i,f(n-i)*i) 0<i<n

*/
func main() {
    fmt.Println(integerBreak(2))
    //fmt.Println(integerBreak(3))
    fmt.Println(integerBreak(4))
    fmt.Println(integerBreak(5))
    //fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
    if n==2 {return 1} //  2,3  拆成至少两个整数会比本身小；其他应该都比本身大，不会出现
    if n==3 {return 2}
    return solve(n,make(map[int]int))
}
func solve(n int,cache map[int]int) int {
    if _,ok:=cache[n];ok{return cache[n]}
    maxi := 1
    for i:=1;i<=n;i++ {
        x := solve(n-i,cache)*i
        if x > maxi {maxi=x}
    }
    cache[n]=maxi
    return maxi
}


