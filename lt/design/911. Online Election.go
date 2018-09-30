package main

import "fmt"

/*
In an election, the i-th vote was cast for persons[i] at time times[i].

Now, we would like to implement the following query function: TopVotedCandidate.q(int t) will return the number of the person that was leading the election at time t.

Votes cast at time t will count towards our query.  In the case of a tie, the most recent vote (among tied candidates) wins.



Example 1:

Input: ["TopVotedCandidate","q","q","q","q","q","q"], [[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
Output: [null,0,1,1,0,0,1]
Explanation:
At time 3, the votes are [0], and 0 is leading.
At time 12, the votes are [0,1,1], and 1 is leading.
At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
This continues for 3 more queries at time 15, 24, and 8.


Note:

1 <= persons.length = times.length <= 5000
0 <= persons[i] <= persons.length
times is a strictly increasing array with all elements in [0, 10^9].
TopVotedCandidate.q is called at most 10000 times per test case.
TopVotedCandidate.q(int t) is always called with t >= times[0].
*/

func main() {
    obj := Constructor([]int{0,1,1,0,0,1,0},[]int{0,5,10,15,20,25,30})
    //[null,0,1,1,0,0,1]
    fmt.Println(obj.Q(3))
    fmt.Println(obj.Q(12))
    fmt.Println(obj.Q(25))
    fmt.Println(obj.Q(15))
    fmt.Println(obj.Q(24))
    fmt.Println(obj.Q(8))
    //*/
}
type Vote struct {
    time int // time
    max int // top vote persion
}
type TopVotedCandidate struct {
    // 时间有序保持最大值
    times []Vote
}
/*
给两个数组，persion是候选人，times是时间序列。
涉及Q函数，传入时间，给出当前vote最高票数。
如果票数相等个，则用最近的一个。所以直接的map是不行的。
*/

func Constructor(persons []int, times []int) TopVotedCandidate {
    cnt := make(map[int]int) // key is persion, value is numbers
    tms := make([]Vote,0)
    max,maxp := 0,0
    // times是有序的。
    for i:=0;i<len(persons);i++ {
        cnt[persons[i]] += 1
        if cnt[persons[i]] >= max {
            // 维护当前最大值
            max,maxp = cnt[persons[i]],persons[i]
        }
        tms = append(tms,Vote{times[i],maxp})
    }
    return TopVotedCandidate{tms}
}


func (this *TopVotedCandidate) Q(t int) int {
    /*
    找第一个小于等于target的数字
    */
    l,r := 0,len(this.times)
    for l<r {
        mid := (l+r)/2
        if this.times[mid].time<=t {
            l=mid+1
        } else {
            r=mid
        }
    }
    return this.times[r-1].max
}



/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */