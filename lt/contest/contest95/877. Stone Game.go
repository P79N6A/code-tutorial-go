package main

import "fmt"

func main() {
        fmt.Println(stoneGame([]int{4,3,4,3,2,5}))
        fmt.Println(stoneGame([]int{7,8,8,10}))
        fmt.Println(stoneGame([]int{3,7,2,3}))
        fmt.Println(stoneGame([]int{7,7,12,16,41,48,41,48,11,9,34,2,44,30,27,12,11,39,31,8,23,11,47,25,15,23,4,17,11,50,16,50,38,34,48,27,16,24,22,48,50,10,26,27,9,43,13,42,46,24}))
}

func stoneGame(piles []int) bool {
        //return canWin(1,0,0,piles,0,len(piles),make(map[int]bool))
        return solve2(piles,0,len(piles)-1)>=0
}
//谁都去得是当前的最大收益
func solve2(nums []int, s,e int) int {
        if s == e {
                return  nums[s]
        }
        a := nums[s] - solve2(nums,s+1,e)
        b := nums[e] - solve2(nums,s,e-1)
        if a > b {
                //fmt.Println(a,nums[s])
                return a
        }
        //fmt.Println(b,nums[e])
        return b
}
type Key struct {
        who,left,right,sum1,sum2 int
}
func canWin(who,sum1,sum2 int, piles[]int,left,right int,cache map[int]bool) bool {
        if left>=right {
                if who == 1{
                        return sum1>sum2
                }
                return sum2>sum1
        }
        if right-left==1 {
                if who==1 {
                        return sum1+piles[left]>sum2
                }
                return sum2+piles[left]>sum1
        }
        key := left << 22 + right << 25 + who << 20 + sum1 << 10 + sum2
        //key := Key{who,0,0,sum1,sum2}
        if _,ok := cache[key];ok {return cache[key]}
        if who==1{
                if !canWin(2,sum1+piles[left],sum2,piles,left+1,right,cache) {
                        cache[key]=true
                        return true
                }
                if !canWin(2,sum1+piles[right-1],sum2,piles,left,right-1,cache){
                        cache[key]=true
                        return true
                }
        } else {
                if !canWin(1,sum1,sum2+piles[left],piles,left+1,right,cache) {
                        cache[key]=true
                        return true
                }
                if !canWin(1,sum1,sum2+piles[right-1],piles,left,right-1,cache) {
                        cache[key]=true
                        return true
                }
        }
        cache[key]=false
        return false
}







