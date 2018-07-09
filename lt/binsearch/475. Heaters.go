package main

import "sort"

/*
Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.

Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.

So, your input will be the positions of houses and heaters seperately, and your expected output will be the minimum radius standard of heaters.

Note:
Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
As long as a house is in the heaters' warm radius range, it can be warmed.
All the heaters follow your radius standard and the warm radius will the same.
Example 1:
Input: [1,2,3],[2]
Output: 1
Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
Example 2:
Input: [1,2,3,4],[1,4]
Output: 1
Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.

*/
func findRadius(houses []int, heaters []int) int {
    max := 0
    sort.Ints(heaters)
    for _,h := range houses {
        idx := binarysearch(heaters,h)
        if idx == 0 {
            if max < heaters[idx]-h {
                max = heaters[idx]-h
            }
        } else if idx > len(heaters)-1 {
            if max < h-heaters[len(heaters)-1] {
                max = h-heaters[len(heaters)-1]
            }
        } else {
            min := heaters[idx]-h
            if min > h-heaters[idx-1] {
                min = h-heaters[idx-1]
            }
            if max < min {
                max = min
            }
        }
    }
    return max
}

func binarysearch(data []int,target int) int {
    l,r := 0,len(data)
    for l < r {
        m := (l+r)/2
        if data[m] < target {
            l = m+1
        } else {
            r = m
        }
    }
    return r
}

func main() {
    findRadius()
    
}
