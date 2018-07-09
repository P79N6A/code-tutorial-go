package cornercase

import (
    "sort"
    "fmt"
)

/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.

*/

/**
 * Definition for an interval.
*/
type Interval struct {
    Start int
    End   int
}
type ISort []Interval
func (s ISort)Len() int {
    return len(s)
}
func (s ISort)Less(i,j int) bool {
    return s[i].Start < s[j].Start
}
func (s ISort)Swap(i,j int) {
    s[i],s[j]=s[j],s[i]
}

func merge(intervals []Interval) []Interval {
    if len(intervals)<=0 {return intervals}
    start,end := intervals[0].Start,intervals[0].End
    ret := make([]Interval,0)
    for i:=1;i<len(intervals);i++ {
        if end > intervals[i].Start {
            if end < intervals[i].End {
                end = intervals[i].End
            }
        } else {
            ret = append(ret,Interval{start,end})
            start,end = intervals[i].Start,intervals[i].End
        }
    }
    ret = append(ret,Interval{start,end})
    return ret
}
func merge1(intervals []Interval) []Interval {
    if len(intervals)<=0 {return intervals}
    sort.Sort(ISort(intervals))
    ret := make([]Interval,0)
    start,end := intervals[0].Start,intervals[0].End
    for i:=1;i<len(intervals);i++ {
        if intervals[i].Start <= end {
            if end < intervals[i].End {
                end = intervals[i].End
            }
        } else {
            ret = append(ret,Interval{start,end})
            start = intervals[i].Start
            end = intervals[i].End
        }
    }
    ret = append(ret,Interval{start,end})
    return ret
}

func main() {
    inter := []Interval{
        Interval{1,3},
        Interval{3,6},
        Interval{8,10},
        Interval{15,18},
    }
    fmt.Println(merge(inter))
}
