package main

import "fmt"
/*
Given a positive integer N, how many ways can we write it as a sum of consecutive positive integers?

Example 1:

Input: 5
Output: 2
Explanation: 5 = 5 = 2 + 3
Example 2:

Input: 9
Output: 3
Explanation: 9 = 9 = 4 + 5 = 2 + 3 + 4
Example 3:

Input: 15
Output: 4
Explanation: 15 = 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
*/
func consecutiveNumbersSum(N int) int {
        /*
        找到规律....
        N%1 = 1
        N%2 = 1
        N%3 = 0
        N%4 = 2
        N%5 = 0
        N%6 = 3
        */
        counter := 1
        for i:=2;2*N>=i*i;i++ {
                if i % 2 != 0 {
                        if N%i == 0 {
                                counter += 1
                        }
                } else {
                        if N%i == i/2 {
                                counter += 1
                        }
                }
        }
        return counter
}
func consecutiveNumbersSum1(N int) int {
        sum := make(map[int]bool)
        s := 0
        sum[0]=true
        for i:=1;i<=N;i++ {
                s += i
                sum[s] = true
        }
        counter := 0
        for k,_ := range sum {
                if _,ok := sum[k-N];ok {
                        counter += 1
                }
        }
        return counter
}

func main() {
        fmt.Println(consecutiveNumbersSum(3))
        fmt.Println(consecutiveNumbersSum(5))
        fmt.Println(consecutiveNumbersSum(9))
        fmt.Println(consecutiveNumbersSum(15))
}
