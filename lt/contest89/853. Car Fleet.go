package main

import (
        "sort"
        "fmt"
)
/*
N cars are going to the same destination along a one lane road.  The destination is target miles away.

Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.

A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.

The distance between these two cars is ignored - they are assumed to have the same position.

A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.


How many car fleets will arrive at the destination?



Example 1:

Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 and 8 become a fleet, meeting each other at 12.
The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
The cars starting at 5 and 3 become a fleet, meeting each other at 6.
Note that no other cars meet these fleets before the destination, so the answer is 3.

Note:

0 <= N <= 10 ^ 4
0 < target <= 10 ^ 6
0 < speed[i] <= 10 ^ 6
0 <= position[i] < target
All initial positions are different.
*/
type pair struct {
        pos int
        speed int
        cost float64
}
type sortPair []pair
func (s sortPair)Len() int {
        return len(s)
}
func (s sortPair)Swap(i,j int) {
        s[i],s[j] = s[j],s[i]
}
func (s sortPair)Less(i,j int) bool {
        return s[i].pos > s[j].pos
}
func carFleet(target int, position []int, speed []int) int {
        ps := make([]pair,0)
        for i:=0;i<len(position);i++ {
                ps = append(ps,pair{position[i],speed[i],float64(target-position[i])/float64(speed[i])})
        }
        sort.Sort(sortPair(ps))
        var currentmax float64 = -1
        num := 0
        for i:=0;i<len(ps);i++ {
                if currentmax==-1 || currentmax<ps[i].cost {
                        currentmax=ps[i].cost
                        num += 1
                }
        }
        return num
}

func main() {
        fmt.Println(carFleet(12,[]int{10,8,0,5,3},[]int{2,4,1,1,3}))

}
