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

type Worker struct {
        factor float64
        q int
}
type sortWorker []Worker
func (s sortWorker)Len() int          {return len(s)}
func (s sortWorker)Less(i,j int) bool {return s[i].factor < s[j].factor}
func (s sortWorker)Swap(i,j int)      {s[i],s[j]=s[j],s[i]}
/////////////////////
type pqInt []int
func (p pqInt)Len() int {return len(p)}
func (p pqInt)Swap(i,j int) {p[i],p[j]=p[j],p[i]}
func (p pqInt)Less(i,j int) bool {return p[i]>p[j]} // 大顶堆
func (p *pqInt)Push(h interface{}) {*p = append(*p,h.(int))}
func (p *pqInt)Pop() interface{} {
        h := (*p)[len(*p)-1]
        *p = (*p)[:len(*p)-1]
        return h
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
        /*
        minimize result R := worker.factor * SumK(worker.q)
        */
        workers := make([]Worker,0)
        for i:=0;i<len(quality);i++ {
                workers = append(workers,Worker{
                        factor:float64(wage[i])/float64(quality[i]),
                        q:quality[i],
                })
        }
        /*
        根据factor[w/q]排序后，后边的factor总能保证前边的满足最低工资付给标准。
        排序后，之后的逻辑处理只跟workers有关了，参数quality和wage都不需要了
        排序保证>k之后待选集合中factor都可用。
         */
        sort.Sort(sortWorker(workers))
        var (
                minMoney float64
                sumQ int
        )
        pq := &pqInt{}
        heap.Init(pq)
        for i:=0;i<K;i++ {
                sumQ += workers[i].q
                heap.Push(pq,int(workers[i].q))
        }
        //先用前K个Q加和，并乘以第K个的worker.factor，计算得到最低付工资的一个值
        // 由于minMoney由worker.factor和sumQ决定，虽然遍历之后的worker.factor在递增，但sumQ可能变小。
        // 为了保证K的长度，每次将之前的最大的Q去掉，将worker.q添加进去。这时就需要大顶堆了。
        // 大顶堆里面维护当前最大的worker.q
        minMoney = float64(sumQ)*workers[K-1].factor
        // priority queue
        for i:=K;i<len(workers);i++ {
                t := heap.Pop(pq) //
                sumQ += workers[i].q-t.(int)
                if minMoney > float64(sumQ) * workers[i].factor {
                        minMoney = float64(sumQ) * workers[i].factor
                }
                heap.Push(pq,int(workers[i].q))
        }
        return minMoney
}


func main() {
        //fmt.Println(mincostToHireWorkers([]int{10,20,5},[]int{70,50,30},2))
        fmt.Println(mincostToHireWorkers([]int{1,2},[]int{14,16},1)) // 14
}
