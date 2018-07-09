package main

import "fmt"
/*
In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

Example 1:

Input: [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
Example 2:

Input: [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat, the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
Note:

1 <= seats.length <= 20000
seats contains only 0s or 1s, at least one 0, and at least one 1.
*/
func maxDistToClosest(seats []int) int {
        maxdis := 0
        first,last := 0,len(seats)-1
        for first < len(seats) {
                if seats[first] == 0 {
                        first += 1
                } else {
                        break
                }
        }
        if first > maxdis {
                maxdis = first
        }
        for last >= 0 {
                if seats[last] == 0 {
                        last -= 1
                } else {
                        break
                }
        }
        if (len(seats)-last-1) > maxdis {
                maxdis = len(seats)-last-1
        }
        idx := first
        for idx < last {
                if seats[idx] == 1 {
                        idx += 1
                        continue
                }
                i := idx
                for ;i<last;i++ {
                        if seats[i]!=0 {
                                break
                        }
                }
                dis := (i-idx+1)/2
                if maxdis < dis {
                        maxdis = dis
                }
                idx = i
        }
        return maxdis
}

func main() {
        fmt.Println(maxDistToClosest([]int{1,0,0,1}))
        fmt.Println(maxDistToClosest([]int{0,0,0,0,1,0,0,0,0,0,1,0,1,0,0,0}))
}
