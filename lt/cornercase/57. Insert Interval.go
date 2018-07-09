package cornercase

import "fmt"

/*
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/
type Interval struct {
    Start int
    End   int
}
/*
不管怎样都需要遍历一次，那么就在遍历过程中处理边界问题等。
*/
func insert(intervals []Interval, newInterval Interval) []Interval {
    ret := make([]Interval,0)
    ins := false

    for i:=0;i<len(intervals);i++ {
        //情况不是非常多。。。。
        if intervals[i].End < newInterval.Start {
            // 不存在交叉情况
            ret = append(ret,intervals[i])
        } else if intervals[i].Start > newInterval.End {
            // 交叉结束
            if !ins {
                ret = append(ret, newInterval)
            }
            ret = append(ret,intervals[i])
            ins = true
        } else {
            // 交叉过程中
            if newInterval.Start > intervals[i].Start {
                newInterval.Start = intervals[i].Start
            }
            if newInterval.End < intervals[i].End {
                newInterval.End = intervals[i].End
            }
        }
    }
    if !ins {
        ret = append(ret,newInterval)
    }
    return ret
}

func main() {
    inter := []Interval{
        Interval{2,5},
        Interval{6,7},
        Interval{8,9},
    }
    fmt.Println(insert(inter,Interval{0,1}))
    /*
    inter := []Interval{
        //Interval{1,3},
        //Interval{6,9},
        Interval{1,2},
        Interval{3,5},
        Interval{6,7},
        Interval{8,10},
        Interval{12,16},
    }
    fmt.Println(insert(inter,Interval{4,8}))
    */
}
