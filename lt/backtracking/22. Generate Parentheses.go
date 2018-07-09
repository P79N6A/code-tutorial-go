package main

import "fmt"
/*

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func generateParenthesis(n int) []string {
        res := make([]string,0)
        solve2(n,n,"",&res)
        return res
}
func solve(left,right int,ret string,res *[]string) {
        if left < 0 || right < 0 || left>right {
                return
        }
        if left == 0 && right == 0 {
                *res = append(*res,ret)
                return
        }
        solve(left-1,right,ret+"(",res)
        solve(left,right-1,ret+")",res)
}
func solve2(left,right int, ret string,res *[]string) {
        if left == 0&&right == 0 {
                *res = append(*res,ret)
        }
        if left >=0 {
                solve2(left-1,right,ret+"(",res)
        }
        if right >left {
                solve2(left,right-1,ret+")",res)
        }
}

func main() {
        fmt.Println(generateParenthesis(1))
        fmt.Println(generateParenthesis(2))
        fmt.Println(generateParenthesis(3))
        fmt.Println(generateParenthesis(4))
        fmt.Println(generateParenthesis(5))
}
