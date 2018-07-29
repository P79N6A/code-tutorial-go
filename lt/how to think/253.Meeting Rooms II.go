package main

import (
    "fmt"
    "sort"
)

/*
描述
给定一系列的会议时间间隔intervals，包括起始和结束时间[[s1,e1],[s2,e2],...] (si < ei)，找到所需的最小的会议室数量。

您在真实的面试中是否遇到过这个题？
样例
给出 intervals = [(0,30),(5,10),(15,20)], 返回 2.
*/
/**
 * Definition of Interval:
*/
type Interval struct {
  Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func minMeetingRooms (intervals []*Interval) int {
    // Write your code here
    start := make([]int,len(intervals))
    end := make([]int,len(intervals))
    for i:=0;i<len(intervals);i++ {
        start[i]=intervals[i].Start
        end[i]=intervals[i].End
    }
    sort.Ints(start)
    sort.Ints(end)
    si,ei := 0,0
    cnt := 0
    max := 0
    for i:=0;i<2*len(intervals);i++ {
        if si < len(intervals) && start[si] < end[ei] {
            si += 1
            cnt += 1
        } else {
            ei += 1
            cnt -= 1
        }
        if cnt > max {
            max = cnt
        }
    }
    return max
}

func main() {
    fmt.Println(minMeetingRooms([]*Interval{
        &Interval{0,30},
        &Interval{5,10},
        &Interval{15,20},
    }))

    
}
