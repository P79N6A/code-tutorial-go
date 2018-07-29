package main

import "fmt"

func main() {
        fmt.Println(minEatingSpeed([]int{3,6,7,11},8))
        fmt.Println(minEatingSpeed([]int{30,11,23,4,2},6))
        fmt.Println(minEatingSpeed([]int{30,11,23,4,2},5))
        fmt.Println(minEatingSpeed([]int{312884470},968709470))
        fmt.Println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184},
        823855818))
}
func minEatingSpeed(piles []int, H int) int {
        sum := 0
        for i:=0;i<len(piles);i++ {
                sum += piles[i]
        }
        l,r := sum/H,sum
        if l <= 0 {l=1}
        for l<r {
                m := l+(r-l)/2
                if eat(piles,m)>H {
                        l=m+1
                } else {
                        r = m
                }
        }
        return r
}
func eat(piles []int,speed int) int {
        n := 0
        for i:=0;i<len(piles);i++ {
                if piles[i]%speed == 0 {
                        n += piles[i]/speed
                } else {
                        n += piles[i]/speed + 1
                }
        }
        return n
}


