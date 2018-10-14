package main

import (
        "sort"
        "fmt"
)
/*
We have jobs: difficulty[i] is the difficulty of the ith job, and profit[i] is the profit of the ith job.

Now we have some workers. worker[i] is the ability of the ith worker, which means that this worker can only complete a job with difficulty at most worker[i].

Every worker can be assigned at most one job, but one job can be completed multiple times.

For example, if 3 people attempt the same job that pays $1, then the total profit will be $3.  If a worker cannot complete any job, his profit is $0.

What is the most profit we can make?

Example 1:

Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
Output: 100
Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get profit of [20,20,30,30] seperately.
Notes:

1 <= difficulty.length = profit.length <= 10000
1 <= worker.length <= 10000
difficulty[i], profit[i], worker[i]  are in range [1, 10^5]
*/

func main() {
        //fmt.Println(maxProfitAssignment([]int{2,4,6,8,10},[]int{10,20,30,40,50},[]int{4,5,6,1}))
        //fmt.Println(maxProfitAssignment([]int{13,37,58},[]int{4,90,96},[]int{34,73,45})) // 190
        // difficulty 和 profit未必是相同顺序的
        fmt.Println(maxProfitAssignment2([]int{68,35,52,47,86},[]int{67,17,1,81,3},[]int{92,10,85,84,82})) // 324
}
/////////////////////////////////////////
func maxProfitAssignment2(difficulty []int, profit []int, worker []int) int {
        // sort + two pointer
        /*
        difficulty 有序, worker有序
        双指针比较两个数组,比较过程保存当前最大profit
        */
        // M = len(difficulty) N = len(worker)
        // O(M*logM+N*logM)
        jobs := make([]*job, 0)
        for i := 0; i < len(difficulty); i++ {
                jobs = append(jobs, &job{difficulty[i], profit[i], profit[i]})
        }
        sort.Sort(joblist(jobs))
        sort.Ints(worker)
        j := 0
        maxprofit := 0
        ans := 0
        for _,w := range worker {
                for j < len(jobs) && w >= jobs[j].difficulty {
                        if maxprofit < jobs[j].profit {
                                maxprofit = jobs[j].profit
                        }
                        j += 1
                }
                ans += maxprofit
        }
        return ans
}


/////////////////////////////////////////
type job struct {
        difficulty,profit int
        curMaxProfit int // 排序后,当前为止最大值
}
type joblist []*job
func (o joblist)Swap(i,j int) {o[i],o[j] = o[j],o[i]}
func (o joblist)Len() int {return len(o)}
func (o joblist)Less(i,j int) bool {
        if o[i].difficulty == o[j].difficulty {
                return o[i].profit < o[j].profit
        }
        return o[i].difficulty < o[j].difficulty
}
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
        // sort + max + binarysearch
        // M = len(difficulty) N = len(worker)
        // 复杂度O(M*logM+N*logM)
        // 两个算法哪个更优需要看M,N的大小.
        jobs := make([]*job,0)
        for i:=0;i<len(difficulty);i++ {
                jobs = append(jobs,&job{difficulty[i],profit[i],profit[i]})
        }
        sort.Sort(joblist(jobs))
        max := jobs[0].profit
        for _,j := range jobs {
                if max < j.profit {
                        max = j.profit
                }
                j.curMaxProfit = max // 当前difficulty的最大profit
        }

        ans := 0
        for _,w := range worker {
                i := upperBound(jobs,w)
                if i >= 0 {
                        ans += jobs[i].curMaxProfit
                }
        }
        return ans
}
func upperBound(jobs []*job,t int) int {
        l,r := 0,len(jobs)
        for l < r {
                m := (l + r)/2
                if jobs[m].difficulty <= t {
                        l = m+1
                } else {
                        r=m
                }
        }
        return r-1
}
///////////////////////////////////////////