package main

import "fmt"

/*
A frog is crossing a river. The river is divided into x units and at each unit there may or may not exist a stone. The frog can jump on a stone, but it must not jump into the water.

Given a list of stones' positions (in units) in sorted ascending order, determine if the frog is able to cross the river by landing on the last stone. Initially, the frog is on the first stone and assume the first jump must be 1 unit.

If the frog's last jump was k units, then its next jump must be either k - 1, k, or k + 1 units. Note that the frog can only jump in the forward direction.

Note:

The number of stones is ≥ 2 and is < 1,100.
Each stone's position will be a non-negative integer < 231.
The first stone's position is always 0.
Example 1:

[0,1,3,5,6,8,12,17]

There are a total of 8 stones.
The first stone at the 0th unit, second stone at the 1st unit,
third stone at the 3rd unit, and so on...
The last stone at the 17th unit.

Return true. The frog can jump to the last stone by jumping
1 unit to the 2nd stone, then 2 units to the 3rd stone, then
2 units to the 4th stone, then 3 units to the 6th stone,
4 units to the 7th stone, and 5 units to the 8th stone.
Example 2:

[0,1,2,3,4,8,9,11]

Return false. There is no way to jump to the last stone as
the gap between the 5th and 6th stone is too large.
*/

func main() {
    fmt.Println(canCross([]int{0,1,3,5,6,8,12,17}))
    fmt.Println(canCross([]int{0,1,2,3,4,8,9,11}))

}
func canCross(stones []int) bool {
    /*
    递归的思路：从头开始，尝试各个pos + (k-1,k,k+1)的组合
    */
    return solve(stones,0,0,make(map[int]bool))
}
func solve(stones []int, pos int, dis int,cache map[int]bool) bool {
    /*
    缓存需要有两个变量，记录从什么位置起跳，跳数是多少；能否到达目的地
    map的key是两个变量，可以借鉴做法是使用bit
    */
    key := pos | dis << 10 // 题目给定pos<1100 < 2^10
    if pos >= len(stones)-1 {return true}
    if _,ok := cache[key];ok{return cache[key]}
    for i:=pos+1;i<len(stones);i++ {
        ndis := stones[i]-stones[pos]
        if ndis > dis+1 {
            // 说明肯定到达不了目的地了，因为数组有序，ndis只会越来越大，提前退出
            cache[key]=false
            return false
        }
        if ndis < dis-1 {continue}
        if solve(stones,i,ndis,cache){
            cache[key]=true
            return true
        }
    }
    cache[key]=false
    return false
}