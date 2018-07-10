package main

import (
        "math"
        "fmt"
)

/*
There are N workers.  The i-th worker has a quality[i] and a minimum wage expectation wage[i].

Now we want to hire exactly K workers to form a paid group.  When hiring a group of K workers, we must pay them according to the following rules:

Every worker in the paid group should be paid in the ratio of their quality compared to other workers in the paid group.
Every worker in the paid group must be paid at least their minimum wage expectation.
Return the least amount of money needed to form a paid group satisfying the above conditions.



Example 1:

Input: quality = [10,20,5], wage = [70,50,30], K = 2
Output: 105.00000
Example 2:

Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], K = 3
Output: 30.66667


Note:

1 <= K <= N <= 10000, where N = quality.length = wage.length
1 <= quality[i] <= 10000
1 <= wage[i] <= 10000
Answers within 10^-5 of the correct answer will be considered correct.


*/


func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
        var min float64 = math.MaxFloat64
        qr := make([]int,0)
        wr := make([]int,0)
        solve(quality,wage,0,K,&qr,&wr,&min)
        return min
}
func caulate(quality []int, wage []int) float64 {
        var minsum float64 = math.MaxFloat64
        for i:=0;i<len(wage);i++ {
                base := float64(wage[i])/float64(quality[i])
                var sum float64
                var valid bool = true
                for i:=0;i<len(quality);i++ {
                        s := base * (float64(quality[i]))
                        if s < float64(wage[i]) {
                                valid = false
                                break
                        }
                        sum += s
                }
                if valid {
                        if sum < minsum {
                                minsum = sum
                        }
                }
        }
        return minsum
}
func solve(q []int,w[]int, cnt, num int, qr *[]int,wr*[]int,min *float64) {
        if len(*qr) < num {
                m := caulate(*qr,*wr)
                if m > *min {
                        return
                }
        }
        if len(*qr) == num {
                m := caulate(*qr,*wr)
                if m < *min {
                        *min = m
                }
                return
        }
        for i:=cnt;i<len(q);i++ {
                *qr = append(*qr,q[i])
                *wr = append(*wr,w[i])
                solve(q,w,i+1,num,qr,wr,min)
                *qr = (*qr)[:(len(*qr)-1)]
                *wr = (*wr)[:(len(*wr)-1)]
        }
}
func main() {
        fmt.Println(mincostToHireWorkers([]int{10,20,5},[]int{70,50,30},2))
        cache := make(map[[]int]bool)
        fmt.Print(cache)
}
