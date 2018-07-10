package main

import (
        "sort"
        "container/heap"
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
Explanation: We pay 70 to 0-th worker and 35 to 2-th worker.
Example 2:

Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], K = 3
Output: 30.66667
Explanation: We pay 4 to 0-th worker, 13.33333 to 2-th and 3-th workers seperately.


Note:

1 <= K <= N <= 10000, where N = quality.length = wage.length
1 <= quality[i] <= 10000
1 <= wage[i] <= 10000
Answers within 10^-5 of the correct answer will be considered correct.
*/

type factor struct {
        f float64
        w float64
        q float64
}
type sortFactor []factor
func (s sortFactor)Len() int {
        return len(s)
}
func (s sortFactor)Less(i,j int) bool {
        return s[i].f < s[j].f
}
func (s sortFactor)Swap(i,j int) {
        s[i],s[j]=s[j],s[i]
}
/////////////////////
type pqInt []int
func (p pqInt)Len() int {
        return len(p)
}
func (p pqInt)Swap(i,j int) {
        p[i],p[j]=p[j],p[i]
}
func (p pqInt)Less(i,j int) bool {
        return p[i]>p[j]
}
func (p *pqInt)Push(h interface{}) {
        *p = append(*p,h.(int))
}
func (p *pqInt)Pop() interface{} {
        h := (*p)[len(*p)-1]
        *p = (*p)[:len(*p)-1]
        return h
}


func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
        /*
        result is : factor * SumK(Q)
        */
        factors := make([]factor,0)
        for i:=0;i<len(quality);i++ {
                factors = append(factors,factor{
                        f:float64(wage[i])/float64(quality[i]),
                        w:float64(wage[i]),
                        q:float64(quality[i]),
                })
        }
        // sort.
        sort.Sort(sortFactor(factors))
        var minMoney,sumQ float64
        pq := &pqInt{}
        heap.Init(pq)
        for i:=0;i<K;i++ {
                sumQ += factors[i].q
                heap.Push(pq,quality[i])
        }
        fmt.Println(factors)
        fmt.Println(pq)
        minMoney = sumQ*factors[K-1].f
        // priority queue
        for i:=K;i<len(quality);i++ {
                t := heap.Pop(pq)
                fmt.Println(t)
                fmt.Println(minMoney,sumQ)
                sumQ += factors[i].q-float64(t.(int))
                fmt.Println(minMoney,sumQ)
                if minMoney > sumQ * factors[i].f {
                        minMoney = sumQ * factors[i].f
                }
                heap.Push(pq,int(factors[i].q))
        }
        return minMoney
}


func main() {
        //fmt.Println(mincostToHireWorkers([]int{10,20,5},[]int{70,50,30},2))
        fmt.Println(mincostToHireWorkers([]int{1,2},[]int{14,16},1)) // 14
}
