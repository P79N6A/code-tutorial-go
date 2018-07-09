package main

import "fmt"

/*

Given a positive integer n and you can do operations as follow:

If n is even, replace n with n/2.
If n is odd, you can replace n with either n + 1 or n - 1.
What is the minimum number of replacements needed for n to become 1?

Example 1:

Input:
8

Output:
3

Explanation:
8 -> 4 -> 2 -> 1
Example 2:

Input:
7

Output:
4

Explanation:
7 -> 8 -> 4 -> 2 -> 1
or
7 -> 6 -> 3 -> 2 -> 1

*/
func integerReplacement(n int) int {
    return solve(n,make(map[int]int))
}
func solve(n int, cache map[int]int) int {
    if n == 1 {
        return 0
    }
    if n%2 == 0 {
        return 1 + solve(n/2,cache)
    }
    if cache[n] > 0 {
        return cache[n]
    }
    n1 := solve(n+1,cache)
    n2 := solve(n-1,cache)
    if n1 > n2 {
        return 1 + n2
    }
    cache[n]=1+n1
    return 1+n1
}
func main() {
    // 1234 => 14
    fmt.Println(integerReplacement(1234))
    
}
