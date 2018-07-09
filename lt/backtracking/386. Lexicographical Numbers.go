package main

import "fmt"

/*
Given an integer n, return 1 - n in lexicographical order.

For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].

Please optimize your algorithm to use less time and space. The input size may be as large as 5,000,000.


*/
func lexicalOrder(n int) []int {
    ret := make([]int,0)
    solve(0,n,&ret)
    return ret
}
func solve(pre int,n int, ret *[]int) {
    for i:=0;i<=9;i++ {
        if pre == 0 && i == 0 {
            continue
        }
        if pre*10+i > n {
            break
        }
        *ret = append(*ret,pre*10+i)
        solve(pre*10+i,n,ret)
    }
}
/*
class Solution {
public:
    vector<int> lexicalOrder(int n) {
        vector<int> res;
        solve(0,n,res)
        return res
    }
    solve(int pre, int n, vector<int> res) {
        for (int i:=0;i<=9;i++) {
            if (pre==0 && i==0) {
                continue;
            }
            if pre*10+i>n {
                break;
            }
            res.push_back(pre*10+i)
            solve(pre*10+i,n,res)
        }
    }
};
*/
func main() {
    fmt.Println(lexicalOrder(10))
    
}
